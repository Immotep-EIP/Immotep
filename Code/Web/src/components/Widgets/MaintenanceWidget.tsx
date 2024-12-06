import React, { useState } from "react";
import { List, Button, Tag } from "antd";
import { CheckOutlined } from "@ant-design/icons";

import {MaintenanceTask} from "@/interfaces/Widgets/Widgets.ts";
import { useTranslation } from "react-i18next";
import style from "./MaintenanceWidget.module.css";

const MaintenanceWidget: React.FC = () => {
    const { t } = useTranslation();

    const [tasks, setTasks] = useState<MaintenanceTask[]>([
        { id: 1, description: "Réparation de la chaudière", priority: "high", completed: false },
        { id: 2, description: "Entretien des jardins", priority: "medium", completed: false },
        { id: 3, description: "Révision de l'ascenseur", priority: "high", completed: false },
        { id: 4, description: "Peinture des couloirs", priority: "low", completed: false },
    ]);

    const markAsCompleted = (id: number) => {
        const updatedTasks = tasks.map((task) =>
            task.id === id ? { ...task, completed: true } : task
        );
        setTasks(updatedTasks);
    };

    const renderPriorityTag = (priority: "high" | "medium" | "low") => {
        const colors = {
            high: "red",
            medium: "orange",
            low: "green",
        };
        return <Tag color={colors[priority]} />;
    };

    return (
        <div className={style.maintenanceWidgetContainer}>
            <h4 className={style.maintenanceWidgetTitle}>{t("widgets.maintenance.title")}</h4>
            <List
                dataSource={tasks}
                renderItem={(task) => (
                    <List.Item className={style.maintenanceWidgetListItem}>
                        <div className={style.maintenanceWidgetTask}>
                            {renderPriorityTag(task.priority)}
                            <span
                                className={
                                    task.completed
                                        ? style.maintenanceWidgetTaskCompleted
                                        : style.maintenanceWidgetTaskPending
                                }
                            >
                                {task.description}
                            </span>
                        </div>
                        {!task.completed && (
                            <Button
                                type="primary"
                                size="small"
                                icon={<CheckOutlined />}
                                onClick={() => markAsCompleted(task.id)}
                            >
                                {t("widgets.maintenance.completeButton")}
                            </Button>
                        )}
                    </List.Item>
                )}
            />
        </div>
    );
};

export default MaintenanceWidget;