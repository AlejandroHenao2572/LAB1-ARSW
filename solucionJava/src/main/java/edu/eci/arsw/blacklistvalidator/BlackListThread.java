package edu.eci.arsw.blacklistvalidator;

import edu.eci.arsw.spamkeywordsdatasource.HostBlacklistsDataSourceFacade;
import java.util.*;
import java.util.concurrent.atomic.AtomicInteger;

class BlackListThread extends Thread {
  
    private int inicio; // rango inicial de hostnames a revisar
    private int fin;    // rango final de hostnames a revisar
    private String ipaddress; // dirección IP a revisar
    private HostBlacklistsDataSourceFacade skds; // fuente de datos de listas negras
    private int ocurrencesCount; // contador de ocurrencias encontradas
    private LinkedList<Integer> blackListOcurrences; // lista de ocurrencias encontradas
    private int checkedServersCount; // contador de servidores revisados
    private AtomicInteger globalOccurrences;

    public BlackListThread(int inicio, int fin, String ipaddress, AtomicInteger globalOccurrences) {
        this.inicio = inicio;
        this.fin = fin;
        this.ipaddress = ipaddress;
        this.skds = HostBlacklistsDataSourceFacade.getInstance();
        this.ocurrencesCount = 0;
        this.blackListOcurrences = new LinkedList<>();
        this.checkedServersCount = 0;
        this.globalOccurrences = globalOccurrences;
    }
    
    @Override
    // Metodo que se ejecuta al iniciar el hilo y busca occurrencias en el rango asignado
    public void run() {
        for (int i = inicio; i < fin; i++) {

            // Verificar si el contador global ha alcanzado el umbral
            if (globalOccurrences.get() >= HostBlackListsValidator.BLACK_LIST_ALARM_COUNT) {
                break; // Salir del bucle si se ha alcanzado el umbral
            }

            checkedServersCount++; // Incrementar el contador por cada servidor revisado
            if (skds.isInBlackListServer(i, ipaddress)) {
                blackListOcurrences.add(i);
                ocurrencesCount++;
                globalOccurrences.incrementAndGet(); // Incrementar el contador global
            }
        }
    }
    
    
    public int getOcurrencesCount() {
        return ocurrencesCount;
    }
    
    public List<Integer> getBlackListOcurrences() {
        return blackListOcurrences;
    }
    
    public int getCheckedServersCount() {
        return checkedServersCount;
    }
}