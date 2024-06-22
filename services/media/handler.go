package media

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"monify/lib"
	"monify/lib/media"
	monify "monify/protobuf/gen/go"
	"net/http"
	"strings"
	"time"
)

type UploadImageResponse struct {
	Url     string `json:"url"`
	ImageId string `json:"image_id"`
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	userId, ok := r.Context().Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse our multipart form, a maximum upload of 5 MB files.
	parseErr := r.ParseMultipartForm(5 << 20)
	if parseErr != nil {
		log.Printf("Could not parse multipart form: %v", parseErr)
		http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
		return
	}

	// Get the usage of the image
	usageId := r.FormValue("usage")
	usage := media.Parse(usageId)
	if usage == monify.Usage_Undefined {
		http.Error(w, "Invalid image usage.", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Error(w, "Could not retrieve the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the image data & type
	imageData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error Reading the File")
		fmt.Println(err)
		http.Error(w, "Could not read the file", http.StatusBadRequest)
		return
	}

	if !strings.Contains(http.DetectContentType(imageData), "image") {
		http.Error(w, "該檔案不是圖片", http.StatusBadRequest)
		return
	}

	imageStorage := r.Context().Value(lib.FileServiceContextKey{}).(media.Storage)
	imgId := uuid.New()
	path := generatePath(usage, fileHeader.Filename, imgId)
	err = imageStorage.Store(path, imageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpFile := media.TmpFile{
		Path:          path,
		ExpectedUsage: usage,
		Uploader:      userId,
		Id:            imgId,
		UploadedAt:    time.Now().UTC(),
	}

	err = StoreTmpFile(r.Context(), tmpFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Success Response
	res := UploadImageResponse{
		Url:     imageStorage.GetUrl(path),
		ImageId: imgId.String(),
	}
	w.Header().Set("Content-Type", "application/json")
	resBody, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func generatePath(usage monify.Usage, originFileName string, imgId uuid.UUID) string {
	now := time.Now()
	fn := usage.String() + "/" + imgId.String() + "." + extractFileNameSuffix(originFileName)
	dir := fmt.Sprintf("%d/%d/%d", now.Year(), now.Month(), now.Day())
	return dir + "/" + fn
}

func extractFileNameSuffix(fileName string) string {
	split := strings.Split(fileName, ".")
	return split[len(split)-1]
}

func StoreTmpFile(ctx context.Context, file media.TmpFile) error {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	_, err := db.ExecContext(ctx, "INSERT INTO tmp_file (fileid, path, uploader, expected_usage, uploaded_at) VALUES ($1, $2, $3, $4, $5)", file.Id, file.Path, file.Uploader, file.ExpectedUsage, file.UploadedAt)
	return err
}
