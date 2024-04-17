# Producer and consumer in Go
- This program aims to write producer and consumer problem in go programming language.
- A shared resource called market is created and shared between producer and consumer threads.
- Cond from sync package is utilized to broadcast messages between threads.
- Two Conds are used, one for producer threads and one for consumer threads. Producer threads will broadcast to consumer threads when market is full and consumer threads will broadcast to producer threads when market is empty.

### Note: 
This will only work in linux. As the clear command used to clear terminal to create the animated output is only for linux.