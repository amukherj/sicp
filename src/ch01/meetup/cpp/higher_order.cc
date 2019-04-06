#include <iostream>
using namespace std;

int main() {
	auto compose = [](auto f, auto g) -> auto {
		return [=](auto arg) auto {
			return f(g(arg));
		};
	};

	auto make_adder = [](auto n) -> auto {
		return [=](auto x) auto {
			return (x + n);
		};
	};

	auto make_multiplier = [](auto n) -> auto {
		return [=](auto x) auto {
			return (x * n);
		};
	};

	cout << compose(make_multiplier(3), make_adder(1))(10)
		 << '\n';
}
