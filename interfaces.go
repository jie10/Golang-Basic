package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type PostgresNumberStore struct {
	// postgres values (db connection)
}

func (p PostgresNumberStore) Put(num int) error {
	//TODO implement me
	fmt.Println("store the number into the Postgres")
	return nil
}

func (p PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

type MongoDBNumberStore struct {
	// some value
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(num int) error {
	//TODO implement me
	fmt.Println("store the number into the MongoDB")
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	// logic
	if err := apiServer.numberStore.Put(1); err != nil {
		panic(err)
	}

	number, err := apiServer.numberStore.GetAll()
	if err != nil {
		return
	}
	fmt.Println(number)
}
