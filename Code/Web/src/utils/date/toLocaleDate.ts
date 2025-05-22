const toLocaleDate = (isoDate: string): string => {
  try {
    const date = new Date(isoDate)

    if (Number.isNaN(date.getTime())) {
      throw new Error('Date invalide')
    }

    const localeToUse = localStorage.getItem('lang') || 'fr'

    const options: Intl.DateTimeFormatOptions = {
      weekday: 'long',
      day: 'numeric',
      month: 'long',
      year: 'numeric'
    }

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
