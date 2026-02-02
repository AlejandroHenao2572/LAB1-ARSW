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
public class CountThreadsMain {
// Ejercio de punto 1     


    public static void main(String a[]){
        
        // Se crean las instancias de los hilos y se le dan los rangos en los que operara cada hilo
        
        CountThread ct1 = new CountThread(0, 99);
        CountThread ct2 = new CountThread(99, 199);
        CountThread ct3 = new CountThread(200, 299);

        // Corre los hilos hilos usando start 
        ct1.start();
        ct2.start();
        ct3.start();
        System.out.println("Threads ended");
        
        // Corre los hilos usando run 
        //ct1.run();
        //ct2.run();
        //ct3.run();

    }
    
}
