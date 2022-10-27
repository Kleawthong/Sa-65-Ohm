import React, { useState,useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { ScholarshipInterface } from "../interfaces/IScholarship";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { GetScholarships } from "../services/HttpClientService";

function Scholarships() {
    const [scholarships, setScholarships] = useState<ScholarshipInterface[]>([]);
    
    useEffect(() => {
      getScholarships();
    }, []);
    const getScholarships = async () => {
      let res = await GetScholarships();
    if (res) {
      setScholarships(res);
    } 
      };
    
   
  
    const columns: GridColDef[] = [
        { field: "ID", headerName: "ID", width: 50 },
        { field: "ScholarName", headerName: "Scholar Name", width: 150 },
        { field: "ScholarAdminID", headerName: "Scholar AdminID", width: 150  },
        { field: "ScholarStatusID", headerName: "Scholar StatusID", width: 150 },
        { field: "ScholarTypeID", headerName: "Scholar TypeID", width: 150 },
        { field: "ScholarDetail", headerName: "Scholar Detail", width: 300 },
      ];
      
   

      return (
        <div>
          <Container maxWidth="md">
            <Box
              display="flex"
              sx={{
                marginTop: 2,
              }}
            >
              <Box flexGrow={1}>
                <Typography
                  component="h2"
                  variant="h6"
                  color="primary"
                  gutterBottom
                >
                  ScholarshipsData
                </Typography>
              </Box>

            </Box>
            <div style={{ height: 400, width: '100%', marginTop: '20px'}}>
              <DataGrid
                rows={scholarships}
                getRowId={(row) => row.ID}
                columns={columns}
                pageSize={5}
                rowsPerPageOptions={[5]}
              />
            </div>
          </Container>
        </div>
      );
            }
      
     export default Scholarships;