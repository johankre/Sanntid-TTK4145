// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

//Mutex? Trådene må ikke bruke cpu samtidig

int i = 0;

pthread_mutex_t counting;

// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    for (int j = 1; j<=1000000; j++) {
        pthread_mutex_lock(&counting);
        i++;
        pthread_mutex_unlock(&counting);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    for (int j = 1; j<=1000000; j++) {
        pthread_mutex_lock(&counting);
        i--;
        pthread_mutex_unlock(&counting);
    }
    return NULL;
}


int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    pthread_mutex_init(&counting, NULL);
    pthread_t thread_1;
    pthread_t thread_2;
    pthread_create(&thread_1, NULL, incrementingThreadFunction, NULL);
    pthread_create(&thread_2, NULL, decrementingThreadFunction, NULL);
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`    
    pthread_join(thread_1, NULL);
    pthread_join(thread_2, NULL);
    
    printf("The magic number is: %d\n", i);
    return 0;
}
