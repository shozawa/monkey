let fizzbuzz = fn(x) {
    if (x > 100) {
        return;
    } else {
        if (x % 15 == 0) {
            puts("FIZZBUZZ");
        } else {
            if (x % 5 == 0) {
                puts("FIZZ");
            } else {
                if (x % 3 == 0) {
                    puts("BUZZ")
                } else {
                    puts(x)
                }
            }
        }
        fizzbuzz(x + 1);
    }
}

fizzbuzz(1);
