package mem

import (
	"testing"
	"time"

	"github.com/mellotonio/desafiogo/app/domain/transfer"
	"github.com/mellotonio/desafiogo/app/infra/utils"
	"github.com/stretchr/testify/assert"
)

type testCaseSingleTransf struct {
	Case          string
	In            transfer.Transfer
	ExpectedValue transfer.Transfer
	ExpectedError error
}

var (
	NewMemTransferRepository = NewTransferRepository(logger)
)

var PassingStoreCasesTransf = []testCaseSingleTransf{
	{
		Case: "Creating a valid transfer",
		In: transfer.Transfer{
			Id:                     utils.GenUUID(),
			Account_origin_id:      "123-456-789",
			Account_destination_id: "879-432-345",
			Amount:                 40000,
			Created_at:             time.Now(),
		},
		ExpectedError: nil,
	},
	{
		Case: "Creating a valid transfer",
		In: transfer.Transfer{
			Id:                     utils.GenUUID(),
			Account_origin_id:      "879-432-345",
			Account_destination_id: "123-456-789",
			Amount:                 20000,
			Created_at:             time.Now(),
		},
		ExpectedError: nil,
	},
}

/*func generateFakeTransfer() []testCaseSingleTransf {
	var fakeTransfers = []testCaseSingleTransf{}
	var fakeIds = []string{"123-456-789", "345-678-654", "534-645-323", "342-537-543", "432-654-865"}

	for i := 0; i < 5; i++ {
		randomTransfer := rand.Intn(40000)
		fakeTransfer := transfer.NewTransfer(fakeIds[i], fakeIds[i+1], randomTransfer)

		fakeTransfers = append(fakeTransfers, *fakeTransfer)
	}

	return fakeTransfers
}*/

func TestTransferRepo(t *testing.T) {
	t.Run("pass", func(t *testing.T) {

		t.Run("transfer.Repository.Store test", func(t *testing.T) {

			for i := range PassingStoreCasesTransf {
				t.Log(PassingStoreCasesTransf[i].Case)
				assert.NoError(t, NewMemTransferRepository.Store(&PassingStoreCasesTransf[i].In))
			}
		})
	})
}
