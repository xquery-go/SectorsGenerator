package config

import "fmt"

type AssetConfig struct {
  Id int64 `sql:"id"`
  TypeId int64 `sql:"type_id"`
  Order []*Roll
  InspirationTables []*InspirationTable
  GroupConfigs []*GroupConfig
  GridConfigs []*GridConfig
}

func (assetConfig *AssetConfig) TableName(assetConfigType string) string {
  return "plan_config_asset"
}

func (assetConfig *AssetConfig) GetId() *int64 {
  return &assetConfig.Id
}

func (assetConfig *AssetConfig) AddPerterbation(perterbation *AssetConfig) *AssetConfig {
  newConfig := new(AssetConfig)
  newConfig.Id = assetConfig.Id
  newConfig.TypeId = assetConfig.TypeId
  newConfig.Order = append(assetConfig.Order, perterbation.Order...)
  newConfig.InspirationTables = StackInspirationTables(assetConfig.InspirationTables, perterbation.InspirationTables)
  newConfig.GroupConfigs = StackGroupConfigs(assetConfig.GroupConfigs, perterbation.GroupConfigs)
  newConfig.GridConfigs = StackGridConfigs(assetConfig.GridConfigs, perterbation.GridConfigs)

  return newConfig
}

func (assetConfig *AssetConfig) Clone() *AssetConfig {
  newConfig := new(AssetConfig)
  newConfig.Id = assetConfig.Id
  newConfig.TypeId = assetConfig.TypeId
  newConfig.Order = make([]*Roll, len(assetConfig.Order))
  copy(newConfig.Order, assetConfig.Order)
  newConfig.InspirationTables = make([]*InspirationTable, len(assetConfig.InspirationTables))
  copy(newConfig.InspirationTables, assetConfig.InspirationTables)
  newConfig.GroupConfigs = make([]*GroupConfig, len(assetConfig.GroupConfigs))
  copy(newConfig.GroupConfigs, assetConfig.GroupConfigs)
  newConfig.GridConfigs = make([]*GridConfig, len(assetConfig.GridConfigs))
  copy(newConfig.GridConfigs, assetConfig.GridConfigs)
  return newConfig
}

func StackAssetConfigs(firstAssetConfigs []*AssetConfig, secondAssetConfigs []*AssetConfig) []*AssetConfig {
  newAssetConfigs := make([]*AssetConfig, len(firstAssetConfigs))
  for i, assetConfig := range firstAssetConfigs {
    newAssetConfigs[i] = assetConfig.Clone()
  }

  for _, perterbationAssetConfig := range secondAssetConfigs {
    assetConfigStacked := false
    for i, newAssetConfig := range newAssetConfigs {
      if newAssetConfig.TypeId == perterbationAssetConfig.TypeId {
        assetConfigStacked = true
        newAssetConfigs[i] = newAssetConfig.AddPerterbation(perterbationAssetConfig)
        break
      }
    }

    if !assetConfigStacked {
      newAssetConfigs = append(newAssetConfigs, perterbationAssetConfig)
    }
  }

  return newAssetConfigs
}

func FetchAssetConfig(manager *ConfigManager, id int64) *AssetConfig {
  assetConfig := new(AssetConfig)
  manager.Client.Fetch(assetConfig, "", id)
  assetConfig.FetchChildren(manager)
  return assetConfig
}

func CreateEmptyConfigAsset(typeId int64) *AssetConfig {
  return &AssetConfig{TypeId: typeId}
}

func (assetConfig *AssetConfig) FetchChildren(manager *ConfigManager) {
  assetConfig.Order = FetchManyRolls(manager, assetConfig.Id, assetConfig.TableName(""), "order")
  assetConfig.InspirationTables = FetchManyInspirationTables(manager, assetConfig.Id, assetConfig.TableName(""), "inspiration_tables")
  assetConfig.GroupConfigs = FetchManyGroupConfigs(manager, assetConfig.Id, assetConfig.TableName(""), "child_configs")
  assetConfig.GridConfigs = FetchManyGridConfigs(manager, assetConfig.Id, assetConfig.TableName(""), "grids")
}

func FetchManyAssetConfigs(manager *ConfigManager, parentId int64, tableName string, valueName string) []*AssetConfig {
  assetConfigs := make([]*AssetConfig, 0)
  assetConfigTableName := new(AssetConfig).TableName("")
  manager.Client.FetchMany(&assetConfigs, parentId, tableName, assetConfigTableName, valueName, "", false)
  for _, assetConfig := range assetConfigs {
    assetConfig.FetchChildren(manager)
  }

  return assetConfigs
}

func (assetConfig *AssetConfig) GetInspirationTable(inspirationTableName string) *InspirationTable {
  for _, inspirationTable := range assetConfig.InspirationTables {
    if inspirationTable.Name == inspirationTableName {
      return inspirationTable
    }
  }


  panic("GetInspirationTable should always return a value!")
}

func (assetConfig *AssetConfig) GetInspirationTableNames() []string {
  tableNames := []string{}
  for _, inspirationTable := range assetConfig.InspirationTables {
    tableNames = append(tableNames, inspirationTable.Name)
  }

  return tableNames
}

func (assetConfig *AssetConfig) GetGroupConfig(groupConfigName string, groupConfigTypeId int64) *GroupConfig {
  for _, groupConfig := range assetConfig.GroupConfigs {
    if groupConfig.Name == groupConfigName && groupConfig.TypeId == groupConfigTypeId {
      return groupConfig
    }
  }

  panic("GetGroupConfig should always return a value!")
}

func (assetConfig *AssetConfig) GetGroupConfigKeys() []*InspirationKey {
  groupConfigKeys := []*InspirationKey{}
  for _, groupConfig := range assetConfig.GroupConfigs {
    key := &InspirationKey{Type: "GroupConfig", Key: groupConfig.Name, Index: groupConfig.TypeId}
    groupConfigKeys = append(groupConfigKeys, key)
  }

  return groupConfigKeys
}

func (assetConfig *AssetConfig) Print(indent int) {
  for i := 0; i < indent; i++ {
    fmt.Print(" ")
  }
  fmt.Print("CONFIG_ASSET:\n")
  for i := 0; i < indent; i++ {
    fmt.Print(" ")
  }
  fmt.Printf("{Id:%d, TypeId:%d, |InspirationTables|:%d, |GroupConfigs|:%d, |GridConfigs|:%d}\n", assetConfig.Id, assetConfig.TypeId, len(assetConfig.InspirationTables), len(assetConfig.GroupConfigs), len(assetConfig.GridConfigs))

  for _, inspirationTable := range assetConfig.InspirationTables {
    inspirationTable.Print(indent+2)
  }
}
