import { useEffect, useState } from "react";
import GetProperties from "@/services/api/Property/GetProperties.ts";
import { GetProperty } from "@/interfaces/Property/Property.tsx";

const useFetchProperties = () => {
    const [properties, setProperties] = useState<GetProperty[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                setLoading(true);
                const res = await GetProperties();
                if (res) {
                    setProperties(res);
                } else {
                    throw new Error("No data received");
                }
            } catch (err: any) {
                console.error("Error fetching properties:", err.message);
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    return { properties, loading, error };
};

export default useFetchProperties;