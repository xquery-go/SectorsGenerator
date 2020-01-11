package config

import "github.com/kyleady/SectorsOfInequity/screamingvortex/utilities"

type Grid struct {
  Name string `sql:"name"`
  Height int `sql:"height"`
  Width int `sql:"width"`
  ConnectionRange int `sql:"connectionRange"`
  PopulationRate float64 `sql:"populationRate"`
  ConnectionRate float64 `sql:"connectionRate"`
  RangeRateMultiplier float64 `sql:"rangeRateMultiplier"`
  SmoothingFactor float64 `sql:"smoothingFactor"`
  WeightedRegions []*WeightedValue
}

func TestGrid() *Grid {
  weightedRegions := []*WeightedValue{
    &WeightedValue{3, 2320},
    &WeightedValue{2, 320},
    &WeightedValue{4, 3499},
  }
  return &Grid{
    "test config",    //Name string `sql:"name"`
    20,               //Height int `sql:"height"`
    20,               //Width int `sql:"width"`
    3,                //ConnectionRange int `sql:"connectionRange"`
    0.75,             //PopulationRate float64 `sql:"populationRate"`
    0.53,             //RangeRateMultiplier float64 `sql:"rangeRateMultiplier"`
    0.51,             //ConnectionRate float64 `sql:"connectionRate"`
    2.0,              //SmoothingFactor float64 `sql:"smoothingFactor"`
    weightedRegions,  //WeightedRegions []*utilities.WeightedValue
  }
}

func (config *Grid) TableName(gridType string) string {
  return "plan_config_grid"
}

func (config *Grid) GetId() *int64 {
  panic("GetId() not implemented. Config should not be editted.")
}

func LoadGridFrom(client utilities.ClientInterface, id int64) *Grid {
  gridConfig := new(Grid)
  client.Fetch(gridConfig, "", id)
  FetchAllWeightedValues(client, &gridConfig.WeightedRegions, WeightedRegionConfigTag(), id)
  return gridConfig
}
