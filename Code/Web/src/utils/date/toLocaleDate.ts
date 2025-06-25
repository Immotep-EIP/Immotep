const toLocaleDate = (isoDate: string, format: string = ''): string => {
  try {
    const date = new Date(isoDate)

    if (Number.isNaN(date.getTime())) {
      throw new Error('Date invalide')
    }

    const localeToUse = localStorage.getItem('lang') || 'fr'

    const options: Intl.DateTimeFormatOptions = (() => {
      switch (format) {
        case 'long':
          return {
            weekday: 'long',
            day: 'numeric',
            month: 'long',
            year: 'numeric'
          }
        case 'short':
          return {
            day: 'numeric',
            month: 'numeric',
            year: 'numeric'
          }
        case 'mid':
          return {
            day: 'numeric',
            month: 'short',
            year: 'numeric'
          }
        default:
          return {
            month: 'long',
            year: 'numeric'
          }
      }
    })()

    let formattedDate = date.toLocaleDateString(localeToUse, options)

    if (localeToUse.startsWith('fr') && date.getDate() === 1) {
      formattedDate = formattedDate.replace(/1 /, '1er ')
    }

    return formattedDate.charAt(0).toUpperCase() + formattedDate.slice(1)
  } catch (error) {
    console.error('Erreur lors du formatage de la date:', error)
    return 'Date invalide'
  }
}

export default toLocaleDate
