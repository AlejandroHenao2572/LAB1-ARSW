/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package edu.eci.arsw.blacklistvalidator;

import java.util.logging.Level;
import java.util.logging.Logger;

/**
 *
 * @author hcadavid
 */
public class Main {
    
    public static void main(String a[]){
        // Deshabilitar logs de INFO
        Logger.getLogger("edu.eci.arsw.spamkeywordsdatasource").setLevel(Level.OFF);
        
        String ipAddress = "202.24.34.55";
        int numCores = Runtime.getRuntime().availableProcessors();
        
        System.out.println("Nucleos: " + numCores + " | IP: " + ipAddress);
        
        // Configuraciones a probar
        int[] numThreads = {1, numCores, numCores * 2, 50, 100, 1000, 100000}; // Numero de cores

        System.out.println("\nHilos\tTiempo (ms)");
        System.out.println("---------------------");
    
        for (int n : numThreads) {
            HostBlackListsValidator validator = new HostBlackListsValidator();
            
            long inicio = System.currentTimeMillis();
            validator.checkHost(ipAddress, n);
            long fin = System.currentTimeMillis();
            
            System.out.println(n + "\t" + (fin - inicio));
        }
    }
}
