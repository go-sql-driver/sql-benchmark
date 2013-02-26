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

    private static void printResult(String name, long ms, int i) {
        System.out.println(
            name + ": " +
            (ms / 1000.0) + "s " +
            "[ " + Math.round(1000.0 * i / ms) + " queries/second ]"
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
            conn = DriverManager.getConnection("jdbc:mysql://localhost/gotest?user=root&password=");
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



        try {
            long startTime, endTime;
            int i, num = 0;
            String str = "";
            PreparedStatement pStmt;

            stmt = conn.createStatement();
            stmt.setPoolable(true);

            stmt.execute("DROP TABLE IF EXISTS test");
            stmt.execute("CREATE TABLE `test` (`number` int(3) NOT NULL, `str` varchar(4) NOT NULL, PRIMARY KEY (`number`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;");

            // Insert data
            pStmt = conn.prepareStatement("INSERT INTO test (`number`, `str`) VALUES(?, ?)");
            for(i=0; i < 100; i++) {
                pStmt.setInt(1, i);
                pStmt.setString(2, "Test");
                pStmt.execute();
            }
            pStmt.close();

            System.out.println("-------------------------------------------------------------");
            System.out.println("   [10000 * Query 100 Rows]");
            System.out.println("-------------------------------------------------------------");

            // SimpleQuery
            startTime = System.currentTimeMillis();
            for(int rep = 0; rep < 10000; rep++) {
                if (stmt.execute("SELECT number, str FROM test")) {
                    rs = stmt.getResultSet();

                    i = 0;
                    while(rs.next()) {
                        num = rs.getInt(1);
                        if(num != i) {
                            throw new Exception("Result didn't match: " + num + "!=" + i);
                        }
                        i++;
                    }

                    if(i != 100) {
                         throw new Exception("Rows count doesn't match: " + i + "=100");
                    }
                } else {
                    throw new Exception("No result");
                }
            }
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult("SimpleQuery", (endTime-startTime), 10000);

            // PreparedQuery
            startTime = System.currentTimeMillis();
            pStmt = conn.prepareStatement("SELECT number, str FROM test");
            for(int rep = 0; rep < 10000; rep++) {
                if (pStmt.execute()) {
                    rs = pStmt.getResultSet();

                    i = 0;
                    while(rs.next()) {
                        num = rs.getInt(1);
                        if(num != i) {
                            throw new Exception("Result didn't match: " + num + "!=" + i);
                        }
                        i++;
                    }

                    if(i != 100) {
                         throw new Exception("Rows count doesn't match: " + i + "=100");
                    }
                } else {
                    throw new Exception("No result");
                }
            }
            pStmt.close();
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult("PreparedQuery", (endTime-startTime), 10000);


            System.out.println();
            System.out.println("-------------------------------------------------------------");
            System.out.println("   [100 * QueryRow] * 1000");
            System.out.println("-------------------------------------------------------------");

            // SimpleQueryRow
            startTime = System.currentTimeMillis();
            for(int rep = 0; rep < 1000; rep++) {
                for(i=0; i < 100; i++) {
                    if (stmt.execute("SELECT * FROM test WHERE number="+i)) {
                        rs = stmt.getResultSet();

                        if(rs.next()) {
                            num = rs.getInt(1);
                            if(num != i) {
                                throw new Exception("Result didn't match: " + num + "!=" + i);
                            }
                            str = rs.getString(2);
                        } else {
                            throw new Exception("No result");
                        }
                    } else {
                        throw new Exception("No result");
                    }
                }
            }
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult("SimpleQueryRow", (endTime-startTime), 100000);

            // PreparedQueryRow
            startTime = System.currentTimeMillis();
            for(int rep = 0; rep < 1000; rep++) {
                pStmt = conn.prepareStatement("SELECT * FROM test WHERE number=?");
                for(i=0; i < 100; i++) {
                    pStmt.setInt(1, i);
                    if (pStmt.execute()) {
                        rs = pStmt.getResultSet();

                        if(rs.next()) {
                            num = rs.getInt(1);
                            if(num != i) {
                                throw new Exception("Result didn't match: " + num + "!=" + i);
                            }
                            str = rs.getString(2);
                        } else {
                            throw new Exception("No result");
                        }
                    } else {
                        throw new Exception("No result");
                    }
                }
                pStmt.close();
            }
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult("PreparedQueryRow", (endTime-startTime), 100000);


            System.out.println();
            System.out.println("-------------------------------------------------------------");
            System.out.println("   [100000 * Exec]");
            System.out.println("-------------------------------------------------------------");

            // SimpleExec
            startTime = System.currentTimeMillis();
            for(i=0; i < 100000; i++) {
                stmt.execute("SET @test_var=1");
            }
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult("SimpleExec", (endTime-startTime), 100000);

            // PreparedExec
            startTime = System.currentTimeMillis();
            pStmt = conn.prepareStatement("SET @test_var=1");
            for(i=0; i < 100000; i++) {
                pStmt.execute("SET @test_var=1");
            }
            pStmt.close();
            endTime = System.currentTimeMillis();
            SQLBenchmark.printResult("PreparedExec", (endTime-startTime), 100000);

        } catch (Exception e) {
            System.err.println(e);
        } finally {
            try {
                stmt.execute("DROP TABLE IF EXISTS test");
            } catch (Exception iDontCare) {}

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
