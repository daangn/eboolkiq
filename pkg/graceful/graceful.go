// Copyright 2021 Danggeun Market Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graceful

import (
	"os"
	"os/signal"
	"syscall"
)

// Stop runs handlers if syscall.SIGINT or syscall.SIGTERM captured.
//
// Stop shall be used so that main function can be terminated.
// This mean defer should not be included in handlers.
//
// For example
//   func main() {
//       lis, err := net.Listen("tcp", "somewhere")
//       if err != nil {
//           log.Fatal(err)
//       }
//       defer lis.Close() // lis.Close should not be included in handlers
//
//       ctx, cancel := context.WithCancel(context.Background())
//
//       // If graceful.Stop capture syscall.SIGINT or syscall.SIGTERM, than run
//       // cancel() so that app.RunInfiniteLoopUntilContextDone will stop
//       graceful.Stop(cancel)
//
//       if err := app.RunInfiniteLoopUntilContextDone(ctx); err != nil {
//           // Since error can occur, defer should not be included in handlers
//           log.Println("something wrong", err)
//       } else {
//           log.Println("server closing")
//       }
//   }
func Stop(handlers ...func()) {
	c := make(chan os.Signal)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func(c chan os.Signal) {
		<-c
		for _, handler := range handlers {
			handler()
		}
		close(c)
	}(c)
}
