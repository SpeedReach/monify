package media

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"monify/lib"
	"monify/lib/media"
	monify "monify/protobuf/gen/go"
)

type Service struct {
	monify.MediaServiceServer
}

func (Service) ConfirmImageUsage(ctx context.Context, req *monify.ConfirmImageUsageRequest) (*monify.MEmpty, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	tmpImage := media.TmpImage{}
	err := db.QueryRowContext(ctx, "SELECT path, expected_usage, uploader, uploaded_at FROM tmpimage WHERE imgid = $1", req.ImageId).Scan(&tmpImage.Path, &tmpImage.ExpectedUsage, &tmpImage.Uploader, &tmpImage.UploadedAt)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "image not found")
	}

	return nil, status.Errorf(codes.Unimplemented, "method ConfirmImageUsage not implemented")
}
