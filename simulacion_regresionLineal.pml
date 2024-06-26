#define N 100

int x[N] = {5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100,
            105, 110, 115, 120, 125, 130, 135, 140, 145, 150, 155, 160, 165, 170, 175, 180,
            185, 190, 195, 200, 205, 210, 215, 220, 225, 230, 235, 240, 245, 250, 255, 260,
            265, 270, 275, 280, 285, 290, 295, 300, 305, 310, 315, 320, 325, 330, 335, 340,
            345, 350, 355, 360, 365, 370, 375, 380, 385, 390, 395, 400, 405, 410, 415, 420,
            425, 430, 435, 440, 445, 450, 455, 460, 465, 470, 475, 480, 485, 490, 495, 500};

int y[N] = {10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180,
            190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340,
            350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480, 490, 500,
            510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640, 650, 660,
            670, 680, 690, 700, 710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820,
            830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940, 950, 960, 970, 980};

int sum_x, sum_y, sum_xy, sum_x_squared, n, m, b;

byte mutex = 1;

active [2] proctype calculate() {
    int i;

    if
    :: mutex == 1 -> 
        atomic {
            mutex = 0;

            n = N;
            sum_x = 0;
            sum_y = 0;
            sum_xy = 0;
            sum_x_squared = 0;

            for (i : 0..n-1) {
                sum_x = sum_x + x[i];
                sum_y = sum_y + y[i];
                sum_xy = sum_xy + (x[i] * y[i]);
                sum_x_squared = sum_x_squared + (x[i] * x[i]);
            }

            m = (n * sum_xy - sum_x * sum_y) / (n * sum_x_squared - sum_x * sum_x);
            b = (sum_y - m * sum_x) / n;

            printf("Process %d - Linear Regression Equation: y = %dx + %d\n", _pid, m, b);

            assert(m > 0);

            mutex = 1;
        }
    :: else ->
        skip;
    fi;
}
