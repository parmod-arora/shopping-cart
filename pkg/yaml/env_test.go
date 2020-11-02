package yaml_test

import (
	"testing"

	"cinemo.com/shoping-cart/pkg/yaml"
	"github.com/google/go-cmp/cmp"
)

func TestFetchEnvVarsFromYaml(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name        string
		args        args
		wantNewEnvs yaml.Envs
		wantErr     bool
	}{
		{
			name: "success load config",
			wantNewEnvs: yaml.Envs{
				Env: map[string]string{
					"SIGNING_PRIVATE_KEY": "PRIVATE KEY\n",
					"SIGNING_PUB_KEY":     "PUBLIC KEY",
				},
			},
			args: args{
				filePath: "testdata/dummy.yml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewEnvs, err := yaml.FetchEnvVarsFromYaml(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchEnvVarsFromYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotNewEnvs, tt.wantNewEnvs) {
				t.Errorf("FetchEnvVarsFromYaml() = %v", cmp.Diff(gotNewEnvs, tt.wantNewEnvs))
			}
		})
	}
}
