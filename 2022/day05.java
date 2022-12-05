import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class day05 {
    public static void main(String[] args) {
        long startTime = System.nanoTime();
        final int STACK_NUMBER = 9;
        ArrayList<Character>[] cratesFirst = new ArrayList[STACK_NUMBER];
        for (int i = 0; i < STACK_NUMBER; i++) {
            cratesFirst[i] = new ArrayList<Character>();
        }
        ArrayList<Character>[] cratesSecond = new ArrayList[STACK_NUMBER];
        for (int i = 0; i < STACK_NUMBER; i++) {
            cratesSecond[i] = new ArrayList<Character>();
        }
                      
        boolean allStackLoaded = false;
        try {
            Scanner myReader = new Scanner(new File("input.txt"));
            while (myReader.hasNextLine()) {  
                String data = myReader.nextLine();
                if(!allStackLoaded){
                    if((int)data.charAt(1) != 49) {
                        int charIndex = 1;
                        for(int i = 0; i < STACK_NUMBER; i++){
                            if((int)data.charAt(charIndex) != 10){
                                cratesFirst[i].add(data.charAt(charIndex));
                                cratesSecond[i].add(data.charAt(charIndex));
                            }
                            charIndex += 4;
                        }
                    }else{                
                        allStackLoaded = true;
                        cratesFirst = reverse(cratesFirst);
                        cratesSecond = reverse(cratesSecond);
                        print(cratesFirst);
                    }
                }else if(data.length() != 0){
                    String[] dati = data.split(" ");
                    int cranesToMove = Integer.parseInt(dati[1]);
                    int startCrane = Integer.parseInt(dati[3]);
                    int endCrane = Integer.parseInt(dati[5]);
                    for(int i = 0; i < cranesToMove; i++){
                        cratesFirst[endCrane-1].add(cratesFirst[startCrane-1].remove(cratesFirst[startCrane-1].size()-1));
                    }
                    for(int i = cranesToMove; i > 0; i--){
                        cratesSecond[endCrane-1].add(cratesSecond[startCrane-1].remove(cratesSecond[startCrane-1].size()-i));
                    }
                }
            }
        } catch (FileNotFoundException e) {
            System.out.println("Invalid input file.");
            e.printStackTrace();
        }
        System.out.print("FIRST PART: ");
        for(int i = 0; i < cratesFirst.length; i++){
            System.out.print(cratesFirst[i].get(cratesFirst[i].size()-1));
        }
        System.out.print("\nSECOND PART: ");
        for(int i = 0; i < cratesSecond.length; i++){
            System.out.print(cratesSecond[i].get(cratesSecond[i].size()-1));
        }
        System.out.println("\nExecution time: "+(System.nanoTime()-startTime)/1000000 + " millis");
    }

    public static ArrayList<Character>[] reverse(ArrayList<Character>[] arr){
        ArrayList<Character>[] reversed = new ArrayList[arr.length];
        for (int i = 0; i < arr.length; i++) {
            reversed[i] = new ArrayList<Character>();
        }
        for (int i = 0; i < arr.length; i++) {
            for (int j = arr[i].size() - 1; j >= 0; j--) {
                reversed[i].add(arr[i].get(j));
            }
        }
        return reversed;
    }

    public static void print(ArrayList<Character>[] crates){
        for(int i = 0; i < crates.length; i++){
            System.out.print("Crate "+(i+1)+" : [");
            for(int j = 0; j < crates[i].size(); j++){
                System.out.print(crates[i].get(j)+" ");
            }
            System.out.println("]");
        }
    } 
}
