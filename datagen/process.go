package datagen

import (
	"github.com/google/uuid"
	"github.com/mbn18/dream/internal/entity"
	"math/rand"
	"strings"
	"time"
)

var (
	pUsers   = []string{"user1", "user2", "user3"}
	commands = []string{"cmd1", "cmd2", "cmd3"}
	args     = []string{"-arg1", "-arg2", "-arg3"}
	tty      = []string{"?", "tty1", "tty2"}
	stat     = []string{"?", "stat1", "stat2"}
)

func genProcessList(executedAt time.Time) entity.ProcessList {
	id := uuid.New()
	num := rand.Intn(4) + 3
	pl := entity.ProcessList{
		CreatedAt: executedAt,
		Processes: make([]*entity.Process, num),
	}
	for i := 0; i < num; i++ {
		pl.Processes[i] = genProcess(id, i+1, executedAt)

	}
	return pl

}

func genProcess(gid uuid.UUID, pid int, executedAt time.Time) *entity.Process {
	return &entity.Process{
		Id:      uuid.New(),
		GroupId: gid,
		User:    pUsers[rand.Intn(len(pUsers))],
		PID:     pid,
		CPU:     rand.Float64() * 100,
		Memory:  rand.Float64() * 100,
		VSZ:     rand.Intn(1000),
		RSS:     rand.Intn(1000),
		TTY:     tty[rand.Intn(len(tty))],
		Stat:    stat[rand.Intn(len(stat))],
		Start:   getRandomTimeOfLast12Hours(executedAt),
		CPUTime: time.Minute * time.Duration(rand.Intn(1000)),
		Command: commands[rand.Intn(len(commands))],
		Args:    getRandomArgs(),
	}
}

func getRandomTimeOfLast12Hours(executedAt time.Time) time.Time {
	duration := time.Duration(rand.Intn(int(time.Hour * 12)))
	return executedAt.Add(-duration)
}

func getRandomArgs() string {
	num := rand.Intn(len(args) + 1)
	return strings.Join(args[:num], " ")
}
