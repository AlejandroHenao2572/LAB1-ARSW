/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package edu.eci.arsw.blacklistvalidator;

import edu.eci.arsw.spamkeywordsdatasource.HostBlacklistsDataSourceFacade;
import java.util.LinkedList;
import java.util.List;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.concurrent.atomic.AtomicInteger;

/**
 *
 * @author hcadavid
 */
public class HostBlackListsValidator {

    static final int BLACK_LIST_ALARM_COUNT=5;
    AtomicInteger globalOccurrencesCount = new AtomicInteger(0); //Contador global de ocurrencias, compartido entre hilos
    
    /**
     * Check the given host's IP address in all the available black lists,
     * and report it as NOT Trustworthy when such IP was reported in at least
     * BLACK_LIST_ALARM_COUNT lists, or as Trustworthy in any other case.
     * The search is not exhaustive: When the number of occurrences is equal to
     * BLACK_LIST_ALARM_COUNT, the search is finished, the host reported as
     * NOT Trustworthy, and the list of the five blacklists returned.
     * @param ipaddress suspicious host's IP address.
     * @return  Blacklists numbers where the given host's IP address was found.
     */
    public List<Integer> checkHost(String ipaddress, int n){
        
        LinkedList<Integer> blackListOcurrences=new LinkedList<>();
        
        int ocurrencesCount=0; //Contador de ocurrencias encontradas
        int checkedListsCount=0; // Contador de listas negras revisadas
        
        HostBlacklistsDataSourceFacade skds=HostBlacklistsDataSourceFacade.getInstance();

        //Calcular la cantidad de servidores por hilo
        int Totalserves = skds.getRegisteredServersCount(); //Obtener el total de servidores
        int serverPerThread = Totalserves / n; // Servidores por hilo
        int resto = Totalserves % n; // Calcular el resto
        BlackListThread[] threads = new BlackListThread[n]; //Lista de hilos

        //Crear los hilos distribuyendo el resto uniformemente
        int currentStart = 0;
        for(int i = 0; i < n; i++) {
            // Los hilos resto reciben 1 servidor extra
            int serversForThisThread = serverPerThread + (i < resto ? 1 : 0);
            int start = currentStart;
            int end = start + serversForThisThread;
            
            //Crear y iniciar el hilo
            threads[i] = new BlackListThread(start, end, ipaddress, globalOccurrencesCount);
            threads[i].start();
            
            currentStart = end; // Actualizar el inicio para el próximo hilo
        }

        //Recoger los resultados de los hilos
        for(int i = 0; i < n ; i++) {
            try {
                // Esperar a que el hilo termine
                threads[i].join();
                // Acumular los resultados
                ocurrencesCount += threads[i].getOcurrencesCount();
                checkedListsCount += threads[i].getCheckedServersCount();
                // Añadir las ocurrencias encontradas por el hilo a la lista principal
                blackListOcurrences.addAll(threads[i].getBlackListOcurrences());
            } catch (InterruptedException e) {
                LOG.log(Level.SEVERE, "Thread interrupted", e);
            }
        }
        
        if (ocurrencesCount>=BLACK_LIST_ALARM_COUNT){
            skds.reportAsNotTrustworthy(ipaddress);
        }
        else{
            skds.reportAsTrustworthy(ipaddress);
        }                
        
        LOG.log(Level.INFO, "Checked Black Lists:{0} of {1}", new Object[]{checkedListsCount, skds.getRegisteredServersCount()});
        
        return blackListOcurrences;
    }
    
    
    private static final Logger LOG = Logger.getLogger(HostBlackListsValidator.class.getName());
}
