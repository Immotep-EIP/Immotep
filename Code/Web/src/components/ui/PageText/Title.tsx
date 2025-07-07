import React from 'react'

import style from './Text.module.css'

interface PageTitleProps {
  title: string
  size?: 'title' | 'subtitle'
  margin?: boolean
  id?: string
}

const PageTitle: React.FC<PageTitleProps> = ({
  title,
  size,
  margin = true,
  id
}) => (
  <span
    id={id}
    className={style.pageTitle}
    style={{
      fontSize: size === 'subtitle' ? '1rem' : '1.4rem',
      fontWeight: size === 'subtitle' ? 400 : 500,
      marginBottom: margin ? '.5rem' : 0
    }}
  >
    {title}
  </span>
)

export default PageTitle
