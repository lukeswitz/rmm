package mindmap

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

type scannerInputMock struct {
}

func (s scannerInputMock) ReadLines() ([]string, error) {
	return scannerInputReadLinesFunc()
}

var scannerInputReadLinesFunc func() ([]string, error)

func TestCreateMindMap(t *testing.T) {
	type args struct {
		source InputSource
	}
	// mocking
	scannerMock := scannerInputMock{}
	dataTest1 := `{
		"mx": {
			"edu.mx": {
				"itesm.edu.mx": {},
				"tecmilenio.edu.mx": {}
			},
			"itesm.mx": {
				"admision.itesm.mx": {},
				"admisionprepatec.itesm.mx": {},
				"ags.itesm.mx": {},
				"apps.itesm.mx": {},
				"btec.itesm.mx": {},
				"cdj.itesm.mx": {},
				"cegs.itesm.mx": {},
				"chi.itesm.mx": {},
				"dm.itesm.mx": {},
				"exatec1.itesm.mx": {},
				"lag.itesm.mx": {},
				"mty.itesm.mx": {
					"web8.mty.itesm.mx": {}
				},
				"net.itesm.mx": {},
				"queretaro.itesm.mx": {
					"comunicacionypublicidad.queretaro.itesm.mx": {},
					"identidad.queretaro.itesm.mx": {}
				},
				"ruv.itesm.mx": {},
				"rzn.itesm.mx": {},
				"sal.itesm.mx": {},
				"sistema.itesm.mx": {},
				"sitios.itesm.mx": {},
				"slp.itesm.mx": {},
				"sorteotec.itesm.mx": {},
				"tecreview.itesm.mx": {},
				"zac.itesm.mx": {}
			},
			"tecreview.mx": {}
		},
		"soy": {
			"prepatec.soy": {}
		}
	}`

	var treeTest1 map[string]interface{}
	err := json.Unmarshal([]byte(dataTest1), &treeTest1)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name          string
		args          args
		readLinesFunc func() ([]string, error)
		want          map[string]interface{}
		wantErr       bool
	}{
		{
			name: "Test 1: Correctly parsing domains",
			args: args{
				source: scannerMock,
			},
			readLinesFunc: func() ([]string, error) {
				return []string{
					"chi.itesm.mx",
					"itesm.mx",
					"ags.itesm.mx",
					"slp.itesm.mx",
					"tecreview.mx",
					"rzn.itesm.mx",
					"mty.itesm.mx",
					"web8.mty.itesm.mx",
					"sistema.itesm.mx",
					"sorteotec.itesm.mx",
					"prepatec.soy",
					"zac.itesm.mx",
					"ruv.itesm.mx",
					"itesm.edu.mx",
					"lag.itesm.mx",
					"dm.itesm.mx",
					"cegs.itesm.mx",
					"tecreview.itesm.mx",
					"exatec1.itesm.mx",
					"btec.itesm.mx",
					"tecmilenio.edu.mx",
					"net.itesm.mx",
					"comunicacionypublicidad.queretaro.itesm.mx",
					"apps.itesm.mx",
					"sitios.itesm.mx",
					"admision.itesm.mx",
					"cdj.itesm.mx",
					"queretaro.itesm.mx",
					"identidad.queretaro.itesm.mx",
					"admisionprepatec.itesm.mx",
					"sal.itesm.mx",
				}, nil
			},
			want:    treeTest1,
			wantErr: false,
		},
		{
			name: "Test 2: Error while reading input",
			args: args{
				source: scannerMock,
			},
			readLinesFunc: func() ([]string, error) {
				return nil, errors.New("something went wrong")
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.readLinesFunc != nil {
				scannerInputReadLinesFunc = tt.readLinesFunc
			}
			got, err := CreateMindMap(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMindMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMindMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
