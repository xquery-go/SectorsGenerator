package config

import "fmt"

type InspirationTable struct {
  Id int64 `sql:"id"`
  Name string `sql:"name"`
  CountRolls []*Roll
  ConstituentParts []*InspirationTable
  WeightedInspirations []*WeightedValue
  ExtraInspirations []*WeightedValue
}

func (inspirationTable *InspirationTable) TableName(inspirationTableType string) string {
  return "plan_inspiration_table"
}

func (inspirationTable *InspirationTable) GetId() *int64 {
  return &inspirationTable.Id
}

func (inspirationTable *InspirationTable) AddPerterbation(perterbationInspirationTable *InspirationTable) *InspirationTable {
  newInspirationTable := new(InspirationTable)
  newInspirationTable.Id = inspirationTable.Id
  newInspirationTable.Name = inspirationTable.Name
  newInspirationTable.CountRolls = append(inspirationTable.CountRolls, perterbationInspirationTable.CountRolls...)
  newInspirationTable.WeightedInspirations = StackWeightedInspirations(inspirationTable.WeightedInspirations, perterbationInspirationTable.WeightedInspirations)
  newInspirationTable.ConstituentParts = append(inspirationTable.ConstituentParts, perterbationInspirationTable.ConstituentParts...)
  newInspirationTable.ExtraInspirations = StackWeightedInspirations(inspirationTable.ExtraInspirations, perterbationInspirationTable.ExtraInspirations)
  return newInspirationTable
}

func (inspirationTable *InspirationTable) Clone() *InspirationTable {
  newInspirationTable := new(InspirationTable)
  newInspirationTable.Id = inspirationTable.Id
  newInspirationTable.Name = inspirationTable.Name
  newInspirationTable.CountRolls = make([]*Roll, len(inspirationTable.CountRolls))
  copy(newInspirationTable.CountRolls, inspirationTable.CountRolls)
  newInspirationTable.WeightedInspirations = make([]*WeightedValue, len(inspirationTable.WeightedInspirations))
  copy(newInspirationTable.WeightedInspirations, inspirationTable.WeightedInspirations)
  newInspirationTable.ConstituentParts = make([]*InspirationTable, len(inspirationTable.ConstituentParts))
  copy(newInspirationTable.ConstituentParts, inspirationTable.ConstituentParts)
  newInspirationTable.ExtraInspirations = make([]*WeightedValue, len(inspirationTable.ExtraInspirations))
  copy(newInspirationTable.ExtraInspirations, inspirationTable.ExtraInspirations)
  return newInspirationTable
}

func (inspirationTable *InspirationTable) FetchChildren(manager *ConfigManager) {
  inspirationTable.CountRolls = FetchManyRolls(manager, inspirationTable.Id, inspirationTable.TableName(""), "count")
  inspirationTable.WeightedInspirations = FetchManyWeightedInspirations(manager, inspirationTable.Id, inspirationTable.TableName(""), "weighted_inspirations")
  inspirationTable.ExtraInspirations = FetchManyWeightedInspirations(manager, inspirationTable.Id, inspirationTable.TableName(""), "extra_inspirations")
  inspirationTable.ConstituentParts = []*InspirationTable{inspirationTable}
}

func StackInspirationTables(firstInspirationTables []*InspirationTable, secondInspirationTables []*InspirationTable) []*InspirationTable {
  newInspirationTables := make([]*InspirationTable, len(firstInspirationTables))
  for i, firstInspirationTable := range firstInspirationTables {
    newInspirationTables[i] = firstInspirationTable.Clone()
  }

  for _, secondInspirationTable := range secondInspirationTables {
    inspirationTableStacked := false
    for i, newInspirationTable := range newInspirationTables {
      if newInspirationTable.Name == secondInspirationTable.Name {
        inspirationTableStacked = true
        newInspirationTables[i] = newInspirationTable.AddPerterbation(secondInspirationTable)
        break
      }
    }

    if !inspirationTableStacked {
      newInspirationTables = append(newInspirationTables, secondInspirationTable.Clone())
    }
  }

  return newInspirationTables
}

func FetchManyInspirationTables(manager *ConfigManager, parentId int64, tableName string, valueName string) []*InspirationTable {
  inspirationTables := make([]*InspirationTable, 0)
  inspirationTableTableName := new(InspirationTable).TableName("")
  manager.Client.FetchMany(&inspirationTables, parentId, tableName, inspirationTableTableName, valueName, "", false)
  for _, inspirationTable := range inspirationTables {
    inspirationTable.FetchChildren(manager)
  }

  return inspirationTables
}

func (inspirationTable *InspirationTable) RollOnce(perterbation *Perterbation) []int64 {
  return RollWeightedValues(inspirationTable.WeightedInspirations, perterbation)
}


func (inspirationTable *InspirationTable) RollCount(perterbation *Perterbation) int {
  return RollAll(inspirationTable.CountRolls, perterbation)
}

func (inspirationTable *InspirationTable) Print(indent int) {
  for i := 0; i < indent; i++ {
    fmt.Print(" ")
  }
  fmt.Print("INSPIRATION_TABLE:\n")
  for i := 0; i < indent; i++ {
    fmt.Print(" ")
  }
  fmt.Printf("{Id:%d, Name:\"%s\", |CountRolls|:%d, |ConstituentParts|:%d, |WeightedInspirations|:%d, |ExtraInspirations|:%d}\n",
    inspirationTable.Id,
    inspirationTable.Name,
    len(inspirationTable.CountRolls),
    len(inspirationTable.ConstituentParts),
    len(inspirationTable.WeightedInspirations),
    len(inspirationTable.ExtraInspirations),
  )
}
