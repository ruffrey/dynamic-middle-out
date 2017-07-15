# Dynamic Middle-Out reservoir neural network

A custom reservoir-like neural network based off my understanding
of neuroscience and vague reading of high level papers on liquid
state machines (though they are few and far between). May not
be very similar as the math is beyond me at this point, but the
goal is not to be a traditional reservoir computing network.

Key differences as I understand them:

- use of tiny int8 for speed, instead of floats, which fundamentally
is different from other neural networks
- more parts of the system are dynamic (?)
