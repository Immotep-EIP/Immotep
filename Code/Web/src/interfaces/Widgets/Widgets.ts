import React from "react";

export interface Widget {
    i: string;
    children: React.ReactNode;
    name: string;
    logo?: React.ReactElement;
    x: number;
    y: number;
    w: number;
    h: number;
    minW?: number;
    maxW?: number;
    minH?: number;
    maxH?: number;
}

export interface MaintenanceTask {
    id: number;
    description: string;
    priority: "high" | "medium" | "low";
    completed: boolean;
}