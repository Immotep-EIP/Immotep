import React from 'react'
import { Tabs, TabsProps } from 'antd'
import { useTranslation } from 'react-i18next'

import style from './DetailsPage.module.css'

const DetailsPage: React.FC = () => {
  const { t } = useTranslation();

  const onChange = (key: string) => {
    console.log(key)
  }

  const items: TabsProps['items'] = [
    {
      key: '1',
      label: t('components.button.about'),
      children: 'Content of Tab Pane 1',
    },
    {
      key: '2',
      label: t('components.button.damage'),
      children: 'Content of Tab Pane 2',
    },
    {
      key: '3',
      label: t('components.button.inventory'),
      children: 'Content of Tab Pane 3',
    },
    {
      key: '4',
      label: t('components.button.documents'),
      children: 'Content of Tab Pane 4',
    },
  ];

  return (
    <div className={style.mainContainer}>
      <div className={style.firstPart}>
        ok
      </div>
      <Tabs
        defaultActiveKey="1"
        items={items}
        onChange={onChange}
        style={{ width: '100%' }}
      />
    </div>
  );
}

export default DetailsPage
