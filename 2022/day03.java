import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class day03 {

    public static void main(String[] args) {
        long startTime = System.nanoTime();
        int[] occurences1 = new int[52];
        int[] occurences2 = new int[52];
        int[] occurences3 = new int[52];
        int totalCounter = 0;
        int turnCounter = 0;
        try {
            Scanner myReader = new Scanner(new File("input.txt"));
            while (myReader.hasNextLine()) {  
                if(turnCounter == 3){
                    for(int i = 0; i < occurences1.length; i++){
                        if(occurences1[i] >= 1 && occurences2[i] >= 1 && occurences3[i] >= 1) totalCounter += 1 + i;
                    }
                    occurences1 = new int[52];
                    occurences2 = new int[52];
                    occurences3 = new int[52];
                    turnCounter = 0;
                }
                String data = myReader.nextLine();
                for(int i = 0; i < data.length(); i++){
                    if(data.charAt(i)>90 && turnCounter == 0){
                        occurences1[(int)data.charAt(i)-'a']++;
                    }else if(turnCounter == 0){
                        occurences1[(int)data.charAt(i)-'A'+26]++;
                    }else if(data.charAt(i)>90 && turnCounter == 1){
                        occurences2[(int)data.charAt(i)-'a']++;
                    }else if(turnCounter == 1){
                        occurences2[(int)data.charAt(i)-'A'+26]++;
                    }else if(data.charAt(i)>90 && turnCounter == 2){
                        occurences3[(int)data.charAt(i)-'a']++;
                    }else if(turnCounter == 2){
                        occurences3[(int)data.charAt(i)-'A'+26]++;
                    }
                }
                turnCounter++; 
            }
        } catch (FileNotFoundException e) {
            System.out.println("Invalid input file.");
            e.printStackTrace();
        }
        System.out.println("SECOND PART: "+totalCounter);
        System.out.println("Execution time: "+(System.nanoTime()-startTime)/1000000 + " millis");
    }
}
