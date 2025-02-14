import React from 'react'
import { Skeleton } from 'antd'

import style from '@/views/RealProperty/details/tabs/2InventoryTab.module.css'

const CardInventoryLoader: React.FC<{ cards: number }> = ({ cards }) => (
  <div className={style.roomsContainer}>
    {Array(cards)
      .fill(0)
      .map((_, index) => (
        // eslint-disable-next-line react/no-array-index-key
        <div key={index} className={style.roomContainer}>
          <div className={style.roomHeader}>
            <Skeleton.Input size="small" active />
            <Skeleton.Avatar
              active
              shape="circle"
              size={20}
              className={style.picture}
            />
          </div>
          <div className={style.stuffsContainer}>
            {Array(2)
              .fill(0)
              .map((_, index) => (
                // eslint-disable-next-line react/no-array-index-key
                <div key={index} className={style.stuffCard}>
                  <Skeleton.Input size="small" active />
                </div>
              ))}
            <div className={style.stuffCardAdd} role="button" tabIndex={0}>
              <Skeleton.Input size="small" active />
            </div>
          </div>
        </div>
      ))}
  </div>
)

export default CardInventoryLoader
