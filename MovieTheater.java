import java.io.*;
import java.util.*;

public class MovieTheater{

    int totalTickets=200;
    LinkedHashMap<String, Integer> tickets
            = new LinkedHashMap<String, Integer>();

    public void seats(){
        tickets.put("J",20);
        tickets.put("I",20);
        tickets.put("H",20);
        tickets.put("G",20);
        tickets.put("F",20);
        tickets.put("E",20);
        tickets.put("D",20);
        tickets.put("C",20);
        tickets.put("B",20);
        tickets.put("A",15);
    }

    public String seatsAllocation(String ticketRequest){

        String n[]=ticketRequest.split(" ");
        int numOfTickets= Integer.parseInt(n[1]);
        if(numOfTickets>totalTickets)
            return "seats are not available";
        String seatNumbers="";
        for (Map.Entry mapElement : tickets.entrySet())
        {
            String k = (String)mapElement.getKey();
            int v = (int)mapElement.getValue();
            if(v>0){
                if(numOfTickets<=v){
                    for(int i=0;i<numOfTickets;i++){
                        seatNumbers=seatNumbers+k+""+v+',';
                        v=v-1;
                        totalTickets--;
                    }
                    tickets.put(k,v);
                    break;
                }
                else if(k=="A" && numOfTickets<totalTickets){
                    for (Map.Entry innerMapElement : tickets.entrySet())
                    {
                        String k1 = (String)innerMapElement.getKey();
                        int v1 = (int)innerMapElement.getValue();
                        while(v1!=0 && numOfTickets!=0){
                            seatNumbers=seatNumbers+k1+""+v1+",";
                            v1=v1-1;
                            totalTickets--;
                            numOfTickets -=1;
                        }
                        tickets.put(k1,v1);
                    }
                }
            }
        }
        return seatNumbers;
    }

    public static void main(String[] args) {
        MovieTheater w=new MovieTheater();
        // instantiating the seats
        w.seats();
        File inputFile=null;
        if (0 < args.length) {
            inputFile= new File("C:\\Users\\Apuroop\\Desktop\\input.txt");
        }
        BufferedReader br = null;
        try{
            File file= new File("C:\\Users\\Apuroop\\Desktop\\output.txt");
            // creating output file and writing result into it
            file.createNewFile();
            // Creating buffer writer to write the output to the output file
            BufferedWriter writer = new BufferedWriter(new FileWriter(file, true));
            String currentLine;
            // Creating buffer reader to read the input of the input file
            br = new BufferedReader(new FileReader(inputFile));
            // printing the path of the output file
            System.out.println( "output file location " +file.getAbsolutePath());
            while ((currentLine = br.readLine()) != null) {
                String res=w.seatsAllocation(currentLine);
                // writing output into the file
                writer.append(res);
                System.out.println(currentLine+" "+res);
                writer.append("\n");
            }
            writer.close();
        }
        catch(Exception e){
            e.getStackTrace();
        }
    }
}