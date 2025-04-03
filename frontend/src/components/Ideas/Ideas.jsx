import React, { useState, useEffect } from 'react';
import styles from "./style.module.scss";

export default function Ideas (){
    const [data, setData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
  
    useEffect(() => {
        // Function to fetch data
        const fetchData = async () => {
          try {
            setLoading(true);
            const response = await fetch('http://localhost:8080/ideas');
            const result = await response.json();
            
            setData(result);
          } catch (error) {
            setError(error);
          } finally {
            setLoading(false);
          }
        };
    
        fetchData();
      }, []);
    

      if (loading) return <p>Loading...</p>;
      if (error) return <p>Error: {error.message}</p>;
      


      return (
        <div>
          <h1>Fetched Data</h1>
            {
            data.map((idea, i) => (
              <div key={i}>
                <h3>Idea {i + 1}:</h3>
                <p>{idea.description}</p>
                <p><strong>Category:</strong> {idea.category}</p>
              </div>
            ))
            }
        </div>
      );


};