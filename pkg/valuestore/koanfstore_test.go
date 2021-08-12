package valuestore

import (
	"path"
	"testing"

	"github.com/ccremer/greposync/core"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKoanfValueStore_ImplementsInterface(t *testing.T) {
	assert.Implements(t, (*core.ValueStore)(nil), new(KoanfValueStore))
}

func TestKoanfValueStore_loadAndMergeConfig(t *testing.T) {
	tests := map[string]struct {
		expectedConf  map[string]interface{}
		givenSyncFile string
	}{
		"GivenExistingFile_ThenLoadYaml": {
			givenSyncFile: "defaults.yml",
			expectedConf: map[string]interface{}{
				"object": map[string]interface{}{
					"key": "value",
				},
				"array": []interface{}{
					"string",
				},
			},
		},
		"GivenNonParseableFile_ThenIgnore": {
			givenSyncFile: "defaults.ini",
			expectedConf:  map[string]interface{}{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewValueStore()
			result, err := s.loadAndMergeConfig(path.Join("testdata", tt.givenSyncFile))
			require.NoError(t, err)
			assert.Equal(t, tt.expectedConf, result.Raw())
		})
	}
}

func TestKoanfValueStore_loadDataForTemplate(t *testing.T) {
	tests := map[string]struct {
		expectedConf          core.Values
		givenSyncFile         string
		givenTemplateFileName string
	}{
		"GivenExistingSimpleFile_ThenLoadYaml": {
			givenSyncFile:         "sync.yml",
			givenTemplateFileName: "README.md",
			expectedConf: core.Values{
				"title": "Hello World",
				"key":   "value",
			},
		},
		"GivenFileWithDirectoryDefaults_ThenLoadYaml": {
			givenSyncFile:         "advanced.yml",
			givenTemplateFileName: ".github/workflows/release.yaml",
			expectedConf: core.Values{
				"title": "Hello World",
				"key":   "specific",
				"array": "overridden",
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewValueStore()
			k := koanf.New(".")
			require.NoError(t, k.Load(file.Provider(path.Join("testdata", tt.givenSyncFile)), yaml.Parser()))
			result, err := s.loadDataForTemplate(k, tt.givenTemplateFileName)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedConf, result)
		})
	}
}
