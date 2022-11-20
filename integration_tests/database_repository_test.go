package integration_tests

import (
	"context"
	"flag"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
	"github.com/KnightHacks/knighthacks_sponsors/repository"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"reflect"
	"testing"
	"time"
)

var integrationTest = flag.Bool("integration", false, "whether to run integration tests")
var databaseUri = flag.String("postgres-uri", "postgresql://postgres:test@localhost:5432/postgres", "postgres uri for running integration tests")

func TestDatabaseRepository_CreateSponsor(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx   context.Context
		input *model.NewSponsor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Sponsor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			got, err := r.CreateSponsor(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSponsor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSponsor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_DeleteSponsor(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			got, err := r.DeleteSponsor(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteSponsor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteSponsor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetSponsor(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Sponsor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			got, err := r.GetSponsor(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSponsor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSponsor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetSponsors(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx   context.Context
		first int
		after string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Sponsor
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			got, got1, err := r.GetSponsors(tt.args.ctx, tt.args.first, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSponsors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSponsors() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetSponsors() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabaseRepository_UpdateDesc(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx         context.Context
		id          string
		sponsorDesc string
		tx          pgx.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			if err := r.UpdateDesc(tt.args.ctx, tt.args.id, tt.args.sponsorDesc, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDesc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateLogo(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx         context.Context
		id          string
		sponsorLogo string
		tx          pgx.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			if err := r.UpdateLogo(tt.args.ctx, tt.args.id, tt.args.sponsorLogo, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLogo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateName(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx         context.Context
		id          string
		sponsorName string
		tx          pgx.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			if err := r.UpdateName(tt.args.ctx, tt.args.id, tt.args.sponsorName, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateSince(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx          context.Context
		id           string
		sponsorSince time.Time
		tx           pgx.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			if err := r.UpdateSince(tt.args.ctx, tt.args.id, tt.args.sponsorSince, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSince() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateSponsor(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx   context.Context
		id    string
		input *model.UpdatedSponsor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Sponsor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			got, err := r.UpdateSponsor(tt.args.ctx, tt.args.id, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateSponsor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateSponsor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_UpdateTier(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx         context.Context
		id          string
		sponsorTier model.SubscriptionTier
		tx          pgx.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			if err := r.UpdateTier(tt.args.ctx, tt.args.id, tt.args.sponsorTier, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateWebsite(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx         context.Context
		id          string
		sponsorSite string
		tx          pgx.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			if err := r.UpdateWebsite(tt.args.ctx, tt.args.id, tt.args.sponsorSite, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateWebsite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_getSponsorWithQueryable(t *testing.T) {
	type fields struct {
		DatabasePool *pgxpool.Pool
	}
	type args struct {
		ctx       context.Context
		id        string
		queryable database.Queryable
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Sponsor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository.DatabaseRepository{
				DatabasePool: tt.fields.DatabasePool,
			}
			got, err := r.GetSponsorWithQueryable(tt.args.ctx, tt.args.id, tt.args.queryable)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSponsorWithQueryable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSponsorWithQueryable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDatabaseRepository(t *testing.T) {
	if *integrationTest == false {
		t.Skipf("skipping integration test")
	}

	type args struct {
		databasePool *pgxpool.Pool
	}

	pool, err := database.ConnectWithRetries(*databaseUri)
	if err != nil {
		t.Error("unable to connect to database", err)
	}
	tests := []struct {
		name string
		args args
		want *repository.DatabaseRepository
	}{
		{
			name: "default",
			args: args{databasePool: pool},
			want: &repository.DatabaseRepository{DatabasePool: pool},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.NewDatabaseRepository(tt.args.databasePool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabaseRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
