import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class day04 {
    public static void main(String[] args) {
        long startTime = System.nanoTime();
        int firstPartCounter = 0, secondPartCounter = 0;
        try {
            Scanner myReader = new Scanner(new File("input.txt"));
            while (myReader.hasNextLine()) {  
                String[] data = myReader.nextLine().split(",");
                String[] firstRange = data[0].split("-");
                String[] secondRange = data[1].split("-");

                int first = Integer.parseInt(firstRange[0]);
                int second = Integer.parseInt(firstRange[1]);
                int firstS = Integer.parseInt(secondRange[0]);
                int secondS = Integer.parseInt(secondRange[1]);

                if( ((firstS <= second && firstS >= first) && (secondS <= second && secondS >= first)) ||
                    ((first <= secondS && first >= firstS) && (second <= secondS && second >= firstS)) ){
                    firstPartCounter++;
                }

                if( (first >= firstS && first <= secondS) || (second >= firstS && second <= secondS) || 
                    (firstS >= first && firstS <= second) || (secondS >= first && secondS <= second)){
                    secondPartCounter++;
                }
            }
        } catch (FileNotFoundException e) {
            System.out.println("Invalid input file.");
            e.printStackTrace();
        }
        System.out.println("FIRST PART: "+firstPartCounter);
        System.out.println("SECOND PART: "+secondPartCounter);
        System.out.println("Execution time: "+(System.nanoTime()-startTime)/1000000 + " millis");
    }
}
