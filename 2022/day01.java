import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class day01 {
    public static void main(String[] args) {
        long startTime = System.nanoTime();
        int firstE = 0, secondE = 0, thirdE = 0;
        int totalCal = 0;
        try {
            Scanner myReader = new Scanner(new File("input.txt"));
            while (myReader.hasNextLine()) {
                String data = myReader.nextLine();
                if(data.length() == 0){
                    if(totalCal >= firstE) {
                        thirdE = secondE;
                        secondE = firstE;
                        firstE = totalCal;
                    }else if(totalCal >= secondE){
                        thirdE = secondE;
                        secondE = totalCal;
                    }else if(totalCal > thirdE){
                        thirdE = totalCal;
                    }
                    totalCal = 0;
                }else{
                    totalCal += Integer.parseInt(data);
                }
            }
        }catch (FileNotFoundException e) {
            System.out.println("Invalid input file.");
            e.printStackTrace();
        }
        System.out.println("FIRST PART: "+firstE);
        System.out.println("SECOND PART: "+(firstE+secondE+thirdE));
        System.out.println("Execution time: "+(System.nanoTime()-startTime)/1000000 + " millis");
    }
}
