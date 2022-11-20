package integration_tests

import (
	"context"
	"flag"
	"fmt"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_shared/utils"
	"github.com/KnightHacks/knighthacks_sponsors/graph/model"
	"github.com/KnightHacks/knighthacks_sponsors/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"reflect"
	"testing"
	"time"
)

var integrationTest = flag.Bool("integration", false, "whether to run integration tests")
var databaseUri = flag.String("postgres-uri", "postgresql://postgres:test@localhost:5432/postgres", "postgres uri for running integration tests")

var databaseRepository *repository.DatabaseRepository

func TestMain(t *testing.M) {
	flag.Parse()
	// check if integration testing is disabled
	if *integrationTest == false {
		return
	}

	// connect to database
	var err error
	pool, err := database.ConnectWithRetries(*databaseUri)
	if err != nil {
		fmt.Printf("unable to connect to database err=%v\n", err)
		os.Exit(-1)
	}

	databaseRepository = repository.NewDatabaseRepository(pool)
	os.Exit(t.Run())
}

func TestDatabaseRepository_CreateSponsor(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *model.NewSponsor
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Sponsor
		wantErr bool
	}{
		{
			name: "create Netflix",
			args: args{
				ctx: context.Background(),
				input: &model.NewSponsor{
					Name:        "Netflix",
					Tier:        model.SubscriptionTierPlatinum,
					Since:       utils.Ptr(time.Date(1999, 10, 10, 0, 0, 0, 0, time.UTC)),
					Description: utils.Ptr("movies and stuff"),
					Website:     utils.Ptr("netflix.com"),
					Logo:        nil,
				},
			},
			want: &model.Sponsor{
				Name:        "Netflix",
				Tier:        model.SubscriptionTierPlatinum,
				Since:       time.Date(1999, 10, 10, 0, 0, 0, 0, time.UTC),
				Description: utils.Ptr("movies and stuff"),
				Website:     utils.Ptr("netflix.com"),
				Logo:        nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := databaseRepository.CreateSponsor(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSponsor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name != tt.want.Name || got.Since != tt.want.Since || !reflect.DeepEqual(got.Logo, tt.want.Logo) || !reflect.DeepEqual(got.Description, tt.want.Description) || !reflect.DeepEqual(got.Website, tt.want.Website) {
				t.Errorf("CreateSponsor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_DeleteSponsor(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := databaseRepository.DeleteSponsor(tt.args.ctx, tt.args.id)
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

func TestDatabaseRepository_GetSponsors(t *testing.T) {
	type args struct {
		ctx   context.Context
		first int
		after string
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.Sponsor
		filter  *model.SponsorFilter
		want1   int
		wantErr bool
	}{
		{
			name: "get 5 sponsors",
			args: args{
				ctx:   context.Background(),
				first: 5,
				after: "2",
			},
			filter: nil,
			want: []*model.Sponsor{
				{
					ID:          "3",
					Name:        "Microsoft",
					Tier:        model.SubscriptionTierPlatinum,
					Since:       time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
					Description: utils.Ptr("does stuff"),
					Website:     utils.Ptr("microsoft.com"),
				},
				{
					ID:          "4",
					Name:        "Apple",
					Tier:        model.SubscriptionTierGold,
					Since:       time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
					Description: utils.Ptr("does stuff"),
					Website:     utils.Ptr("apple.com"),
				},
				{
					ID:          "5",
					Name:        "Bing",
					Tier:        model.SubscriptionTierPlatinum,
					Since:       time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
					Description: utils.Ptr("does stuff"),
					Website:     utils.Ptr("bing.com"),
				},
				{
					ID:          "6",
					Name:        "Oracle",
					Tier:        model.SubscriptionTierBronze,
					Since:       time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
					Description: utils.Ptr("does stuff"),
					Website:     utils.Ptr("oracle.com"),
				},
				{
					ID:          "7",
					Name:        "UrMom",
					Tier:        model.SubscriptionTierSilver,
					Since:       time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
					Description: utils.Ptr("does stuff"),
					Website:     utils.Ptr("urmom.com"),
				},
			},
			want1:   -1,
			wantErr: false,
		},
		{
			name: "get 2 sponsors",
			args: args{
				ctx:   context.Background(),
				first: 2,
				after: "2",
			},
			filter: &model.SponsorFilter{Tiers: []model.SubscriptionTier{model.SubscriptionTierPlatinum}},
			want: []*model.Sponsor{
				{
					ID:          "3",
					Name:        "Microsoft",
					Tier:        model.SubscriptionTierPlatinum,
					Since:       time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
					Description: utils.Ptr("does stuff"),
					Website:     utils.Ptr("microsoft.com"),
				},
				{
					ID:          "5",
					Name:        "Bing",
					Tier:        model.SubscriptionTierPlatinum,
					Since:       time.Date(2000, 10, 10, 0, 0, 0, 0, time.UTC),
					Description: utils.Ptr("does stuff"),
					Website:     utils.Ptr("bing.com"),
				},
			},
			want1:   2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := databaseRepository.GetSponsors(tt.args.ctx, tt.filter, tt.args.first, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSponsors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want1 != -1 && got1 != tt.want1 {
				t.Errorf("GetSponsors() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSponsors() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_UpdateDesc(t *testing.T) {
	type args struct {
		ctx         context.Context
		id          string
		sponsorDesc string
		queryable   database.Queryable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update Joe Shmoe's description to ",
			args: args{
				ctx:         context.Background(),
				id:          "2",
				sponsorDesc: "fixes the cracks in your wood",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		{
			name: "update invalid sponsor",
			args: args{
				ctx:         context.Background(),
				id:          "-1",
				sponsorDesc: "nah g",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		{
			name: "update invalid sponsor 2",
			args: args{
				ctx:         context.Background(),
				id:          "1253434",
				sponsorDesc: "nah g",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := databaseRepository.UpdateDesc(tt.args.ctx, tt.args.id, tt.args.sponsorDesc, tt.args.queryable); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDesc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateLogo(t *testing.T) {
	type args struct {
		ctx         context.Context
		id          string
		sponsorLogo string
		queryable   database.Queryable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update Joe Shmoe's logo to a picture of wood",
			args: args{
				ctx:         context.Background(),
				id:          "2",
				sponsorLogo: "https://imagesvc.meredithcorp.io/v3/mm/image?q=60&c=sc&poi=face&w=825&h=413&url=https%3A%2F%2Fstatic.onecms.io%2Fwp-content%2Fuploads%2Fsites%2F49%2F2017%2F03%2F28%2F100447495.jpg",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		{
			name: "update invalid sponsor",
			args: args{
				ctx:         context.Background(),
				id:          "-1",
				sponsorLogo: "nah g",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		{
			name: "update invalid sponsor 2",
			args: args{
				ctx:         context.Background(),
				id:          "1253434",
				sponsorLogo: "nah g",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := databaseRepository.UpdateLogo(tt.args.ctx, tt.args.id, tt.args.sponsorLogo, tt.args.queryable); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLogo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateName(t *testing.T) {
	type args struct {
		ctx         context.Context
		id          string
		sponsorName string
		queryable   database.Queryable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update 'Joe Shmoe's Woodworking' to 'Joe Shmoe's Wood Working'",
			args: args{
				ctx:         context.Background(),
				id:          "2",
				sponsorName: "Joe Shmoe's Wood Working",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		{
			name: "update invalid sponsor",
			args: args{
				ctx:         context.Background(),
				id:          "-1",
				sponsorName: "nah g",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		{
			name: "update invalid sponsor 2",
			args: args{
				ctx:         context.Background(),
				id:          "1253434",
				sponsorName: "nah g",
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := databaseRepository.UpdateName(tt.args.ctx, tt.args.id, tt.args.sponsorName, tt.args.queryable); (err != nil) != tt.wantErr {
				t.Errorf("UpdateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateSince(t *testing.T) {
	type args struct {
		ctx          context.Context
		id           string
		sponsorSince time.Time
		queryable    database.Queryable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update Joe Shmoe's 'since' date to 2022/11/9",
			args: args{
				ctx:          context.Background(),
				id:           "2",
				sponsorSince: time.Date(2022, time.November, 9, 0, 0, 0, 0, time.UTC),
				queryable:    databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		{
			name: "update invalid sponsor",
			args: args{
				ctx:          context.Background(),
				id:           "-1",
				sponsorSince: time.Date(2022, time.November, 9, 0, 0, 0, 0, time.UTC),
				queryable:    databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		{
			name: "update invalid sponsor 2",
			args: args{
				ctx:          context.Background(),
				id:           "1253434",
				sponsorSince: time.Date(2022, time.November, 9, 0, 0, 0, 0, time.UTC),
				queryable:    databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		// TODO: Add test cases.}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := databaseRepository.UpdateSince(tt.args.ctx, tt.args.id, tt.args.sponsorSince, tt.args.queryable); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSince() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateSponsor(t *testing.T) {
	type args struct {
		ctx   context.Context
		id    string
		input *model.UpdatedSponsor
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Sponsor
		wantErr bool
	}{
		{
			name: "update abcdef",
			args: args{
				ctx: context.Background(),
				id:  "8",
				input: &model.UpdatedSponsor{
					Name:        utils.Ptr("ghijklmnop"),
					Tier:        utils.Ptr(model.SubscriptionTierPlatinum),
					Since:       utils.Ptr(time.Date(1999, 11, 11, 0, 0, 0, 0, time.UTC)),
					Description: utils.Ptr("nope"),
					Website:     utils.Ptr("google.com"),
					Logo:        utils.Ptr("nahg.com/img.png"),
				},
			},
			want: &model.Sponsor{
				ID:          "8",
				Name:        "ghijklmnop",
				Tier:        model.SubscriptionTierPlatinum,
				Since:       time.Date(1999, 11, 11, 0, 0, 0, 0, time.UTC),
				Description: utils.Ptr("nope"),
				Website:     utils.Ptr("google.com"),
				Logo:        utils.Ptr("nahg.com/img.png"),
			},
			wantErr: false,
		},
		{
			name: "update invalid",
			args: args{
				ctx: context.Background(),
				id:  "2342342",
				input: &model.UpdatedSponsor{
					Name:        utils.Ptr("ghijklmnop"),
					Tier:        utils.Ptr(model.SubscriptionTierPlatinum),
					Since:       utils.Ptr(time.Date(1999, 11, 11, 0, 0, 0, 0, time.UTC)),
					Description: utils.Ptr("nope"),
					Website:     utils.Ptr("google.com"),
					Logo:        utils.Ptr("nahg.com/img.png"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.UpdateSponsor(tt.args.ctx, tt.args.id, tt.args.input)
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
	type args struct {
		ctx         context.Context
		id          string
		sponsorTier model.SubscriptionTier
		queryable   database.Queryable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update Joe Shmoe's subscription tier to Silver",
			args: args{
				ctx:         context.Background(),
				id:          "2",
				sponsorTier: model.SubscriptionTierSilver,
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		{
			name: "update invalid sponsor",
			args: args{
				ctx:         context.Background(),
				id:          "-1",
				sponsorTier: model.SubscriptionTierSilver,
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		{
			name: "update invalid sponsor 2",
			args: args{
				ctx:         context.Background(),
				id:          "1253434",
				sponsorTier: model.SubscriptionTierSilver,
				queryable:   databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := databaseRepository.UpdateTier(tt.args.ctx, tt.args.id, tt.args.sponsorTier, tt.args.queryable); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateWebsite(t *testing.T) {

	type args struct {
		ctx         context.Context
		id          string
		sponsorSite string
		tx          database.Queryable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update Joe Shmoe's website to joe.mama",
			args: args{
				ctx:         context.Background(),
				id:          "2",
				sponsorSite: "joe.mama",
				tx:          databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := databaseRepository.UpdateWebsite(tt.args.ctx, tt.args.id, tt.args.sponsorSite, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateWebsite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_GetSponsor(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Sponsor
		wantErr bool
	}{
		{
			name: "get billy bob",
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want: &model.Sponsor{
				ID:          "1",
				Name:        "Billy Bob LLC",
				Tier:        model.SubscriptionTierPlatinum,
				Since:       time.Date(2022, 11, 9, 0, 0, 0, 0, time.UTC),
				Description: utils.Ptr("loves coding"),
				Website:     utils.Ptr("billybob.com"),
			},
			wantErr: false,
		},
		{
			name: "get -1 ID",
			args: args{
				ctx: context.Background(),
				id:  "-1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get invalid sponsor",
			args: args{
				ctx: context.Background(),
				id:  "2389472938",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetSponsor(tt.args.ctx, tt.args.id)
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

func TestDatabaseRepository_GetSponsorWithQueryable(t *testing.T) {
	type args struct {
		ctx       context.Context
		id        string
		queryable database.Queryable
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Sponsor
		wantErr bool
	}{
		{
			name: "get billy bob",
			args: args{
				ctx:       context.Background(),
				id:        "1",
				queryable: databaseRepository.DatabasePool,
			},
			want: &model.Sponsor{
				ID:          "1",
				Name:        "Billy Bob LLC",
				Tier:        model.SubscriptionTierPlatinum,
				Since:       time.Date(2022, 11, 9, 0, 0, 0, 0, time.UTC),
				Description: utils.Ptr("loves coding"),
				Website:     utils.Ptr("billybob.com"),
			},
			wantErr: false,
		},
		{
			name: "get -1 ID",
			args: args{
				ctx:       context.Background(),
				id:        "-1",
				queryable: databaseRepository.DatabasePool,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get invalid sponsor",
			args: args{
				ctx:       context.Background(),
				id:        "2389472938",
				queryable: databaseRepository.DatabasePool,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := databaseRepository.GetSponsorWithQueryable(tt.args.ctx, tt.args.id, tt.args.queryable)
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
	type args struct {
		databasePool *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *repository.DatabaseRepository
	}{
		{
			name: "default",
			args: args{databasePool: databaseRepository.DatabasePool},
			want: &repository.DatabaseRepository{DatabasePool: databaseRepository.DatabasePool},
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
