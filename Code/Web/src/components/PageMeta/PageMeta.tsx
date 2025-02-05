import React from 'react'
import { Helmet } from 'react-helmet-async'

const DEFAULT_TITLE_PREFIX = 'Immotep'

interface PageMetaProps {
  title: string
  description?: string
  keywords?: string
}

const PageMeta: React.FC<PageMetaProps> = ({
  title,
  description,
  keywords
}) => {
  const fullTitle = `${title} - ${DEFAULT_TITLE_PREFIX}`
  return (
    <Helmet>
      <title>{fullTitle}</title>
      {description && <meta name="description" content={description} />}
      {keywords && <meta name="keywords" content={keywords} />}
    </Helmet>
  )
}

export default PageMeta
