/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package edu.eci.arsw.threads;

/**
 *
 * @author hcadavid
 */

//la clase extiende de Thread 
public class CountThread extends Thread {

    //Inicializacion de los atributos A y B (rangos)
    public int A;
    public int B;

    public CountThread(int A, int B) {
        this.A = A;
        this.B = B;
    }

    //Al correr el hilo se ejecuta el metodo run
    @Override
    public void run() {
        for (int i = A; i <= B; i++) {
            System.out.println(i);
        }
    }
}
