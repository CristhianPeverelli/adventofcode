import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class day06 {

    static final int CHAR_TO_COUNT_FIRST = 4;
    static final int CHAR_TO_COUNT_SECOND = 14;

    public static void main(String[] args) throws FileNotFoundException {
        long startTime = System.nanoTime();

        int firstPartValue = 0;
        int secondPartValue = 0;  
        Scanner myReader = new Scanner(new File("input.txt"));
        String data = myReader.nextLine();          

        //First part
        int[] alphabetOccFirst = new int[26];
        for(int i = 0; i < CHAR_TO_COUNT_FIRST; i++){
            alphabetOccFirst[data.charAt(i) - 'a'] += 1;
        }
        for(int i = CHAR_TO_COUNT_FIRST; i <= data.length(); i++){
            alphabetOccFirst[data.charAt(i-CHAR_TO_COUNT_FIRST)-'a'] -= 1;
            alphabetOccFirst[data.charAt(i)-'a'] += 1;
            if(check(alphabetOccFirst)){
                firstPartValue = i+1;
                break;
            }
        }

        //Second part
        int[] alphabetOccSecond = new int[26];
        for(int i = 0; i < CHAR_TO_COUNT_SECOND; i++){
            alphabetOccSecond[data.charAt(i) - 'a'] += 1;
        }
        for(int i = CHAR_TO_COUNT_SECOND; i <= data.length(); i++){
            alphabetOccSecond[data.charAt(i-CHAR_TO_COUNT_SECOND)-'a'] -= 1;
            alphabetOccSecond[data.charAt(i)-'a'] += 1;
            if(check(alphabetOccSecond)){
                secondPartValue = i+1;
                break;
            }
        } 
        
        System.out.println("FIRST PART: "+firstPartValue);
        System.out.println("SECOND PART: "+secondPartValue);
        System.out.println("Execution time: "+(System.nanoTime()-startTime)/1000000 + " millis");
    }

    private static boolean check(int[] arr){
        for(int i = 0; i < arr.length; i++){
            if(arr[i] > 1){
                return false;
            }
        }
        return true;
    }

}

 
