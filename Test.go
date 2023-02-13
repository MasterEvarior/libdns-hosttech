package main

import (
	"context"
	"github.com/libdns/libdns"
)

func main() {
	apiKey := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiYWEwMDQzMmQ3ODQwMThmMTEyNjAwNTVkMTNiZTZhZjM2NDJiZWVmNWEyMTcyOWZhMTE2ZTAwOGYzOGQ4MWFkZWRkMjJkNzZkMTk0OTc4YjYiLCJpYXQiOjE2NzYwMjk4MDcuOTE5NzYxLCJuYmYiOjE2NzYwMjk4MDcuOTE5NzY4LCJleHAiOjQ4MDAxNjc0MDcuODg2NjgzLCJzdWIiOiI2NDg3Iiwic2NvcGVzIjpbXX0.SuXDd7usxERPajJ-wFQ4x3VnlOOHSTRkYAFJ0Hb2pp1OHUSQ4Pp2-9qp4mUlBkWo_Ue3n3VGpWUozksABr4XGz6yVI15S3A_yPLYDd3IjVCKeHHWijayvdnY3OnZitqNXaAA-7WGz550NmlACsvUmFCxJ59IpYswaG12tCbtfCnFa8mu7-ymFhyY5YXbH_2xS7qRqPHOVZdezh9-1pk_dZbneIzyrzj4ie3r35BgU8cxhpA1FlOXJF2EFPG4NUQSjMXPigNUhJ56ey9wswtqeooeyUTfKdiQQxIB93HD1vlsNfaZJ8u2yd-R-esytD6Rlaes9uuwfIuo7nHTWiv7YVao_AemAa4UxIk09jJSOqcgzXClv2ra809DAnrRbd7F4bojCpZVcBJRYDsmvmU8eIDQPtC8-C7ERa1gDT3SfNORZ2wSWWK5zDESeRp6cuHfy86WdFN98JaBCA7IOddSaTyR6mN5ubyaNhlGw81d9b5QSUTnO4CrVYHxFbuEmLVvUVBfcR8GJNeGOopRwETtuOk9R65f85WB1EurVr2z0bpNzNt4Krc6q0sX8dr-OHQsAxUxXLvyehPUn8EV57KnFzoI6U_8kXECAKtA8Oljl2BO7kxKSjYiCs-HAVEdsHEoSEsAvQKTmb2DWWBRNXtr47y462OwuTkdIar8nn0jgWw"
	provider := Provider{
		APIToken: apiKey,
	}

	provider.GetRecords(context.Background(), "hmlb.ch")

	provider.DeleteRecords(context.Background(), "hmlb.ch", []libdns.Record{{
		ID: "2499489",
	}, {ID: "2499490"}})
}
