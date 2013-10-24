import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

/**
 *
 * @author Julien Schmidt <mail at julienschmidt.com>
 */
public class SQLBenchmark {

    private static void printResult(String name, long ms, long bytes, int i) {
        System.out.println(
            name + ": " +
            (ms / 1000.0) + "s   \t " +
            Math.round(1000.0 * i / ms) + " queries/second   \t"
            /* Does someone know how to get useable values?
             * If yes, please make a pull-request to fix this! */
            //(bytes/i) + " B/query"
        );
    }

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        try {
            Class.forName("com.mysql.jdbc.Driver").newInstance();
        } catch (Exception ex) {
            System.err.println("Loading MySQL-Driver failed!");
        }

        Connection conn = null;
        try {
            conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/gotest?user=root&password=root");
        } catch (SQLException ex) {
            if(conn != null) {
                try {
                    conn.close();
                } catch(Exception iDontCare) {
                } finally {
                    conn = null;
                }
            }
            System.out.println("SQLException: " + ex.getMessage());
            System.out.println("SQLState: " + ex.getSQLState());
            System.out.println("VendorError: " + ex.getErrorCode());
            System.exit(1);
        }

        Statement stmt  = null;
        ResultSet rs    = null;

        Runtime runtime = Runtime.getRuntime();

        try {
            long startTime, endTime;
            long startTotalMem;
            PreparedStatement pStmt;

            int num;

            stmt = conn.createStatement();
            stmt.setPoolable(true);

            // SimpleExec
            startTotalMem = runtime.totalMemory()-runtime.freeMemory();
            startTime = System.currentTimeMillis();
            for(int i = 0; i < 500000; i++) {
                stmt.execute("DO 1");
            }
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult(
            	"SimpleExec",
            	(endTime-startTime),
            	(runtime.totalMemory()-runtime.freeMemory()-startTotalMem),
            	500000
            );

            // PreparedExec
            startTotalMem = runtime.totalMemory()-runtime.freeMemory();
            startTime = System.currentTimeMillis();
            pStmt = conn.prepareStatement("DO 1");
            for(int i = 0; i < 500000; i++) {
                pStmt.execute();
            }
            pStmt.close();
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult(
            	"PreparedExec",
            	(endTime-startTime),
            	(runtime.totalMemory()-runtime.freeMemory()-startTotalMem),
            	500000
            );

            // SimpleQueryRow
            startTotalMem = runtime.totalMemory()-runtime.freeMemory();
            startTime = System.currentTimeMillis();
            for(int i = 0; i < 500000; i++) {
                if (stmt.execute("SELECT 1")) {
                    rs = stmt.getResultSet();

                    if(rs.next()) {
                        num = rs.getInt(1);
                    } else {
                        throw new Exception("No result");
                    }
                } else {
                    throw new Exception("No result");
                }
            }
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult(
            	"SimpleQueryRow",
            	(endTime-startTime),
            	(runtime.totalMemory()-runtime.freeMemory()-startTotalMem),
            	500000
            );

            // PreparedQueryRow
            startTotalMem = runtime.totalMemory()-runtime.freeMemory();
            startTime = System.currentTimeMillis();
            pStmt = conn.prepareStatement("SELECT 1");
            for(int i = 0; i < 500000; i++) {
                if (pStmt.execute()) {
                    rs = pStmt.getResultSet();

                    if(rs.next()) {
                        num = rs.getInt(1);
                    } else {
                        throw new Exception("No result");
                    }
                } else {
                    throw new Exception("No result");
                }
            }
            pStmt.close();
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult(
            	"PreparedQueryRow",
            	(endTime-startTime),
            	(runtime.totalMemory()-runtime.freeMemory()-startTotalMem),
            	500000
            );

            // PreparedQueryRowParam
            startTotalMem = runtime.totalMemory()-runtime.freeMemory();
            startTime = System.currentTimeMillis();
            pStmt = conn.prepareStatement("SELECT ?");
            for(int i = 0; i < 500000; i++) {
            	pStmt.setInt(1, i);
                if (pStmt.execute()) {
                    rs = pStmt.getResultSet();

                    if(rs.next()) {
                        num = rs.getInt(1);
                    } else {
                        throw new Exception("No result");
                    }
                } else {
                    throw new Exception("No result");
                }
            }
            pStmt.close();
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult(
            	"PreparedQueryRowParam",
            	(endTime-startTime),
            	(runtime.totalMemory()-runtime.freeMemory()-startTotalMem),
            	500000
            );

        } catch (Exception e) {
            System.err.println(e);
        } finally {
            if (rs != null) {
                try {
                    rs.close();
                } catch (SQLException iDontCare) {
                } finally {
                    rs = null;
                }
            }

            if (stmt != null) {
                try {
                    stmt.close();
                } catch (SQLException iDontCare) {
                } finally {
                    stmt = null;
                }
            }

            try {
                conn.close();
            } catch(Exception iDontCare) {
            } finally {
                conn = null;
            }
        }
    }
}
