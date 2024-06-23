package media

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"monify/lib"
	"monify/lib/media"
	monify "monify/protobuf/gen/go"
	"time"
)

type Service struct {
	monify.MediaServiceServer
}

func (Service) ConfirmFileUsage(ctx context.Context, req *monify.ConfirmFileUsageRequest) (*emptypb.Empty, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	tmpImage := media.TmpFile{}
	err := db.QueryRowContext(ctx, "SELECT path, expected_usage, uploader, uploaded_at FROM tmp_file WHERE file_id = $1", req.FileId).Scan(&tmpImage.Path, &tmpImage.ExpectedUsage, &tmpImage.Uploader, &tmpImage.UploadedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "image not found")
		}
		return nil, status.Errorf(codes.Internal, "error getting image: %v", err)
	}

	_, err = db.ExecContext(ctx, "INSERT INTO confirmed_file (file_id, path, \"usage\", uploader, uploaded_at, confirmed_at) VALUES ($1, $2, $3, $4, $5, $6)", req.FileId, tmpImage.Path, tmpImage.ExpectedUsage, tmpImage.Uploader, tmpImage.UploadedAt, time.Now())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting image: %v", err)
	}
	_, _ = db.ExecContext(ctx, "DELETE FROM tmp_file WHERE file_id = $1", req.FileId)
	return &emptypb.Empty{}, nil
}

func (Service) GetFileUrl(ctx context.Context, req *monify.GetFileUrlRequest) (*monify.GetFileUrlResponse, error) {
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	var path string
	err := db.QueryRowContext(ctx, "SELECT path FROM confirmed_file WHERE file_id = $1", req.FileId).Scan(&path)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "image not found")
		}
		return nil, status.Errorf(codes.Internal, "error getting image")
	}

	cfg := ctx.Value(lib.ConfigContextKey{}).(Config)
	return &monify.GetFileUrlResponse{Url: cfg.S3Host + path}, nil
}
