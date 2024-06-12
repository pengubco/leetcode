
## On Mac
### VSCode

### Test
https://google.github.io/googletest/primer.html

brew install llvm
brew install googletest

Add to ~/.zshrc
export PATH="/opt/homebrew/opt/llvm/bin:$PATH
export GTEST_HOME=/opt/homebrew/Cellar/googletest/1.14.0

clang++ -std=c++20 -I $GTEST_HOME/include -L $GTEST_HOME/lib app.cpp -lgtest -lgtest_main -pthread -o app
