package model

func MapAssetInputToAsset(a AssetInput, e Asset) Asset { 
 if a.Description != nil { 
	e.Description = *a.Description 
 }
  return e
}
