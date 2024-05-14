package model

import (
		"github.com/cend-org/duval/internal/database"
		"time"
)

type AssetCollector struct {} 
func (c *AssetCollector) Description(a *string) (r Asset, err error) { 
		err = database.Get(&r, `SELECT * FROM asset WHERE description = ? ORDER BY created_at DESC LIMIT 1`, a)
  return r, err
}

func (c *AssetCollector) _b(){ 
		_ = time.Now()
}

func MapAssetInputToAsset(a AssetInput, e Asset) Asset { 
 if a.Description != nil { 
	e.Description = *a.Description 
 }
  return e
}
