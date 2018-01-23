package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetStickerSet_Success(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetStickerSet, "token")).ParamPresent("name").
		Reply(http.StatusOK).JSON(`{
			"ok": true,
			"result": {
				"name": "persik",
				"title": "persik",
				"contains_masks": false,
				"stickers": [
					{
						"width": 512,
						"height": 512,
						"emoji": "😘",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABR3WSoABKwikyJYauM4BEAAAgI",
							"file_size": 4602,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD7AkAAvQ3UQLI7WdKH3y_gwI",
						"file_size": 17424
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😡",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABOso1kqAARCBHuVjiIpab89AAIC",
							"file_size": 5118,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD7gkAAvQ3UQLQybldiVSlYAI",
						"file_size": 20500
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😖",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABN0nVkqAAQcycnXFGtlAnk6AAIC",
							"file_size": 4994,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD8AkAAvQ3UQIzeEanHFKCrAI",
						"file_size": 18314
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😏",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABPoolkqAAQjnHLFoM88rM42AAIC",
							"file_size": 4584,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD8gkAAvQ3UQI5f78QEIIWjQI",
						"file_size": 16272
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😢",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABP9fVkqAARNfgPcPVdyits5AAIC",
							"file_size": 4342,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD9AkAAvQ3UQJ0r9BUBv0NWAI",
						"file_size": 16498
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😳",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABNHeVkqAARsx48_MjWLK_88AAIC",
							"file_size": 5130,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD9gkAAvQ3UQJyoKxllFmATAI",
						"file_size": 19104
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😡",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABNLd1kqAATrEKf1AAFjozQIPgACAg",
							"file_size": 5436,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD-AkAAvQ3UQKJqoL03Q023gI",
						"file_size": 22166
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😔",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABPqe1kqAAQYdWpIxCUFVJ1AAAIC",
							"file_size": 4082,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD-gkAAvQ3UQL2ByeAx12yZQI",
						"file_size": 14540
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😳",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABMDo1kqAAQu4pf9Mxe0lZQ3AAIC",
							"file_size": 5176,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD_QkAAvQ3UQLGhLN3M03BDQI",
						"file_size": 19276
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😎",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABMllFkqAAR8-ZZDfd1kbn82AAIC",
							"file_size": 5814,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgAD_wkAAvQ3UQKjkyuvaYyQAwI",
						"file_size": 21376
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😠",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABPgn1kqAARuCm3u5tu1UKQ9AAIC",
							"file_size": 4302,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADAQoAAvQ3UQJD9zf1SwIOXQI",
						"file_size": 16418
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😠",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABN5dlkqAASLkiRVhH1vhpZHAAIC",
							"file_size": 5020,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADAwoAAvQ3UQIcFX1R0AXmEQI",
						"file_size": 18762
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "👍",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABM9fFkqAAS2NFcjej59XZM9AAIC",
							"file_size": 4422,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADBQoAAvQ3UQLawzI2JExIzAI",
						"file_size": 16646
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😎",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABMdlFkqAASoJBfydQUmol43AAIC",
							"file_size": 5268,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADBwoAAvQ3UQJBx8BJ3nuqygI",
						"file_size": 21350
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😭",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABOBnVkqAAQ4p4So3W9O9O05AAIC",
							"file_size": 3978,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADCQoAAvQ3UQJ50oKBh7vQ6QI",
						"file_size": 13462
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😋",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABMzolkqAARuzGrxLVLdHH8xAAIC",
							"file_size": 4822,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADCwoAAvQ3UQJ4XvvLulkejAI",
						"file_size": 18016
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😍",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABOIfVkqAASVBPW8ZHrBXYdBAAIC",
							"file_size": 4486,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADDQoAAvQ3UQJRbaiw8f73xQI",
						"file_size": 15976
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "🙈",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABPulFkqAAShdnYDWWfpHA8yAAIC",
							"file_size": 4754,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADEQoAAvQ3UQKiTPRu_TKfjAI",
						"file_size": 17176
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😒",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABNPlFkqAATuyvE1egFsRao3AAIC",
							"file_size": 4618,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADEwoAAvQ3UQLjcxz8bvmqgAI",
						"file_size": 16724
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😌",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABNTlFkqAAR5se_mB3B9g0Q0AAIC",
							"file_size": 4386,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADFQoAAvQ3UQLNhSJGfDygMwI",
						"file_size": 15004
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "👎",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABOxpFkqAAQSHAOPGw4EjMI4AAIC",
							"file_size": 3958,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADFwoAAvQ3UQLiLdK7t5KPTQI",
						"file_size": 13216
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😎",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABPxf1kqAASMYXbPoEohBfQ7AAIC",
							"file_size": 4842,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADGQoAAvQ3UQJEjM2Q4Ed0GgI",
						"file_size": 17644
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😏",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABNUf1kqAASkEcdkA9CK0QABSwACAg",
							"file_size": 4274,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADGwoAAvQ3UQLpIl0BtxzUoQI",
						"file_size": 15814
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😁",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABNQflkqAASjJlCFcRQqAwABQwACAg",
							"file_size": 4804,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADHQoAAvQ3UQKOAAF3CrwW7WwC",
						"file_size": 18470
					},
					{
						"width": 512,
						"height": 512,
						"emoji": "😝",
						"set_name": "persik",
						"thumb": {
							"file_id": "AAQCABNXklkqAATv3r0AARtYfHhtOQACAg",
							"file_size": 4850,
							"width": 128,
							"height": 128
						},
						"file_id": "CAADAgADHwoAAvQ3UQKQKLhs3R5mmwI",
						"file_size": 17570
					}
				]
			}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	model, res, err := client.GetStickerSet("persik").Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestGetStickerSet_Error(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointGetStickerSet, "token")).ParamPresent("name").
		Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	model, res, err := client.GetStickerSet("persik").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestGetStickerSet_Failed(t *testing.T) {
	gock.New(telegraph.BaseURL).Get(fmt.Sprintf(telegraph.EndpointGetStickerSet, "token")).ParamPresent("name").
		Reply(http.StatusBadRequest).JSON(`{
			"ok": false,
			"error_code": 400,
			"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	model, res, err := client.GetStickerSet("persik").Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}
