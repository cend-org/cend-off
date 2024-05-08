package model 


 /* AssetInput */ 



func MapAssetInputToAsset(input AssetInput, existing Asset) Asset { 
 if input.Description != nil { 
	existing.Description = *input.Description 
 } 
 
  return existing 
}