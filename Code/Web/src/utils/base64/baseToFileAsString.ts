function base64ToFileAsString(base64String: string) {
  const matches = base64String.match(/^data:(.+?);base64,(.+)$/)
  let mimeType = ''
  let base64Content = ''

  if (matches) {
    ;[mimeType, base64Content] = [matches[1], matches[2]]
  } else {
    mimeType = 'application/octet-stream'
    base64Content = base64String
  }

  const byteCharacters = atob(base64Content)
  const byteNumbers = new Array(byteCharacters.length)
  for (let i = 0; i < byteCharacters.length; i += 1) {
    byteNumbers[i] = byteCharacters.charCodeAt(i)
  }

  const byteArray = new Uint8Array(byteNumbers)

  const blob = new Blob([byteArray], { type: mimeType })
  const fileUrl = URL.createObjectURL(blob)

  return fileUrl
}

export default base64ToFileAsString
