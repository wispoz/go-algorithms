package queue

import "testing"

func TestQueue_Dequeue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)

	val ,err := q.Dequeue()

	if err == false{
		t.Fatal("Cannot dequeue")
	}
	if val != 1 {
		t.Fatal("Value must be 1")
	}
}

func TestQueue_Dequeue2(t *testing.T) {
	q := Queue{}

	_ , err := q.Dequeue()

	if err != false {
		t.Fatal("Error must be false")
	}

}
func TestQueue_Size(t *testing.T) {

	q := Queue{}
	if q.Size() != 0 {
		t.Fatal("Size must be 0")
	}

	q.Enqueue(1)

	if q.Size() != 1 {
		t.Fatal("Size must be 1")
	}

	q.Dequeue()

	if q.Size() != 0 {
		t.Fatal("Size must be 0")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	if q.Size() != 1 {
		t.Fatal("Enqueue failed")
	}

	val,err := q.Dequeue()

	if err == false {
		t.Fatal("Value must be 1")
	}

	if val != 1 {
		t.Fatal("Value must be 1")
	}
}