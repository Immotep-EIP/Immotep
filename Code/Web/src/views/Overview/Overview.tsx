import React from 'react';
import {Responsive, WidthProvider} from 'react-grid-layout';
import {UserOutlined} from '@ant-design/icons';

import {useTranslation} from "react-i18next";
import PageTitle from "@/components/PageText/Title.tsx";
import UserInfoWidget from "@/components/Widgets/UserInfoWidget.tsx";
import MaintenanceWidget from "@/components/Widgets/MaintenanceWidget.tsx";
import {Widget} from "@/interfaces/Widgets/Widgets.ts";
import style from './Overview.module.css';
import "@/../node_modules/react-grid-layout/css/styles.css"
import "@/../node_modules/react-resizable/css/styles.css"

const ResponsiveGridLayout = WidthProvider(Responsive);

const WidgetTemplate: React.FC<{ name: string; logo?: React.ReactElement; children: React.ReactNode }> =
    ({name, logo, children}) => (
        <div className={style.widgetContainer}>
            <div className={style.widgetHeader}>
                {logo}
            </div>
            <div className={style.widgetContent}>
                {children}
            </div>
        </div>
    );

const Overview: React.FC = () => {
    const {t} = useTranslation()
    const layouts: { lg: Widget[] } = {
        lg: [
            {
                i: "widget1",
                name: "Widget 1",
                logo: <UserOutlined/>,
                x: 0,
                y: 0,
                w: 2,
                h: 2,
                children: <UserInfoWidget/>,
                minW: 2,
                maxW: 3,
                minH: 2,
                maxH: 3
            },
            {
                i: "widget2",
                name: "Maintenance",
                logo: <UserOutlined/>,
                x: 2,
                y: 0,
                w: 3,
                h: 4,
                children: <MaintenanceWidget/>,
                minW: 3,
                maxW: 6,
                minH: 4,
                maxH: 6
            }
        ],
    };

    const draggableHandleClass = `.${style.widgetHeader}`;

    return (
        <div className={style.pageContainer}>
            <div className={style.pageHeader}>
                <PageTitle title={t('pages.overview.title')} size="title"/>
            </div>
            <ResponsiveGridLayout
                className={style.gridLayout}
                layouts={layouts}
                breakpoints={{lg: 1200, md: 996, sm: 768, xs: 480, xxs: 0}}
                cols={{lg: 12, md: 10, sm: 6, xs: 4, xxs: 2}}
                rowHeight={80}
                isResizable
                draggableHandle={draggableHandleClass}
            >
                {layouts.lg.map((widget) => (
                    <div key={widget.i} data-grid={widget}>
                        <WidgetTemplate name={widget.name} logo={widget.logo}>
                            {widget.children}
                        </WidgetTemplate>
                    </div>
                ))}
            </ResponsiveGridLayout>

        </div>
    );
};

export default Overview;