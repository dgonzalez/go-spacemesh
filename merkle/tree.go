package merkle

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/spacemeshos/go-spacemesh/log"
	"github.com/syndtr/goleveldb/leveldb"
)

// A general-purpose merkle tree backed by (k,v) stores
// All (k,v) methods are in user data space and not in tree space.
// Tree space pointers and paths are internal only.
type MerkleTree interface {
	Put(k, v []byte) error                // store user key, value
	Delete(k []byte) error                // delete user value indexed by key
	Get(k []byte) ([]byte, *stack, error) // get user value indexed by key
	GetRootHash() []byte                  // get tree root hash
	GetRootNode() NodeContainer           // get root node

	CloseDataStores() error // call when done w the tree

	Print() string
}

// internal implementation
type merkleTreeImp struct {
	userData *leveldb.DB
	treeData *leveldb.DB
	root     NodeContainer
}

// Creates a new empty merkle tree with the provided paths to user and tree data db files.
// The db files will be created on these pathes if they don't already exist.
// userDataFileName: full local os path and file name for the user data db for this tree
// treeDataFileName: full local os path and file name for the internal tree db store for this tree
func NewEmptyTree(userDataFileName string, treeDataFileName string) (MerkleTree, error) {
	userData, err := leveldb.OpenFile(userDataFileName, nil)
	if err != nil {
		log.Error("Failed to open user db %v", err)
		return nil, err
	}

	treeData, err := leveldb.OpenFile(treeDataFileName, nil)
	if err != nil {
		log.Error("Failed to open tree db %v", err)
		return nil, err
	}

	mt := &merkleTreeImp{
		userData: userData,
		treeData: treeData,
	}

	return mt, nil
}

// Creates a new tree from provided dbs file paths.
// rootHash: tree root hash - used to pull the root from the db
// userDataFileName: full local os path and file name for user data db for this tree
// treeDataFileName: full local os path and file name for the internal tree db store for this tree
func NewTreeFromDb(rootHash []byte, userDataFileName string, treeDataFileName string) (MerkleTree, error) {

	userData, err := leveldb.OpenFile(userDataFileName, nil)
	if err != nil {
		return nil, err
	}

	treeData, err := leveldb.OpenFile(treeDataFileName, nil)
	if err != nil {
		return nil, err
	}

	mt := &merkleTreeImp{
		userData: userData,
		treeData: treeData,
	}

	// load the tree from the db
	data, err := treeData.Get(rootHash, nil)
	if err != nil {
		return nil, err
	}

	root, err := newNodeFromData(data)
	if err != nil {
		return nil, err
	}

	mt.root = root
	return mt, nil
}

func (mt *merkleTreeImp) CloseDataStores() error {

	err := mt.treeData.Close()
	if err != nil && err != leveldb.ErrClosed {
		log.Error("Failed to close tree db %v", err)
		return err
	}

	err = mt.userData.Close()
	if err != nil && err != leveldb.ErrClosed {
		log.Error("Failed to close user db %v", err)
		return err
	}

	return nil
}

// Print the tree to a string
func (mt *merkleTreeImp) Print() string {
	buffer := bytes.Buffer{}
	buffer.WriteString("\n------------\n")
	if mt.root == nil {
		buffer.WriteString("Merkle Tree: Empty tree.\n")
	} else {

		buffer.WriteString(fmt.Sprintf("Merkle tree: root hash <%s>\n", hex.EncodeToString(mt.GetRootHash())[:6]))
		buffer.WriteString(mt.root.print(mt.userData, mt.treeData))
	}
	buffer.WriteString("------------\n")
	return buffer.String()
}