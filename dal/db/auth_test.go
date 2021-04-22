package db

import (
	"testing"
)

func Test_getRoles(t *testing.T) {

	gotRoles, err := getRoles(ctx, 1)
	if err != nil {
		t.Errorf("getRoles() error = %v, wantErr %v", err, nil)
		return
	}
	t.Logf("%+v", gotRoles[0])

}

func Test_getResourceIDs(t *testing.T) {

	gotResourceIDs, err := getResourceIDs(ctx, []uint{1})
	if err != nil {
		t.Errorf("getResourceIDs() error = %v, wantErr %v", err, nil)
		return
	}
	t.Logf("%+v", gotResourceIDs)
}

func TestGetAuthList(t *testing.T) {

	gotAuthList, err := GetAuthList(ctx, 1)
	if err != nil {
		t.Errorf("GetAuthList() error = %v, wantErr %v", err, nil)
		return
	}
	t.Logf("%+v", gotAuthList)
}
