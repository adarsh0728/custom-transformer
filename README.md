# Custom Transformer with Panic

This is a simple User Defined Source Transformer example which receives a message, changes the message event time to time.Now(), increments the message counter, panics on every count which is a multiple of 1000, otherwise returns the message.
