import React from 'react'
import { LoadingOutlined } from '@ant-design/icons';

import {useTranslation} from "react-i18next";
import useFetchProperties from "@/hooks/useEffect/useFetchProperties.ts";
import { useAuth } from '@/context/authContext'
import style from './UserInfoWidget.module.css'

const UserInfoWidget: React.FC = () => {
    const { user } = useAuth()
    const { t } = useTranslation()
    const { properties, loading, error } = useFetchProperties();

    if (loading) {
        return (
            <div>
                <p>{t("generals.loading")}</p>
                <LoadingOutlined />
            </div>
        );
    }

    if (error) {
        return <p>{t("widgets.userInfo.errorFetching")}</p>;
    }

    return (
        <div className={style.layoutContainer}>
            {user ? (
                <div key={user.id}>
                    <p>{t("widgets.userInfo.title")} {user.firstname} !</p>
                    <p>
                        {[
                            t("widgets.userInfo.propertiesNumber"),
                            properties.length,
                            t("widgets.userInfo.realProperties")]
                            .join(' ')
                        }
                    </p>
                </div>
            ) : (
                <p>{t("widgets.userInfo.noUser")}</p>
            )}
        </div>
    )
};

export default UserInfoWidget;
