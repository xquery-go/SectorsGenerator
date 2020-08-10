package config

type Inspiration struct {
  Id int64 `sql:"id"`
  Name string `sql:"name"`
  PerterbationIds []int64
  InspirationRolls []*InspirationTable
  InspirationTables []*InspirationTable
}

func FetchInspiration(manager *ConfigManager, inspirationId int64) *Inspiration {
  inspiration := new(Inspiration)
  inspirationTableName := inspiration.TableName("")
  manager.Client.Fetch(inspiration, "", inspirationId)
  inspiration.InspirationTables = FetchManyInspirationTables(manager, inspirationId, inspirationTableName, "inspiration_tables")
  inspiration.InspirationRolls = FetchManyInspirationTables(manager, inspirationId, inspirationTableName, "roll_groups")
  inspiration.PerterbationIds = FetchManyPerterbationIds(manager, inspirationId, inspirationTableName, "perterbations")
  return inspiration
}

func (inspiration *Inspiration) TableName(inspirationType string) string {
  return "plan_inspiration"
}

func (inspiration *Inspiration) GetId() *int64 {
  return &inspiration.Id
}

func FetchManyInspirationIds(manager *ConfigManager, parentId int64, tableName string, valueName string) []int64 {
  ids := make([]int64, 0)
  exampleInspiration := new(Inspiration)
  manager.Client.FetchManyToManyChildIds(&ids, parentId, tableName, exampleInspiration.TableName(""), valueName, "", false)
  return ids
}
