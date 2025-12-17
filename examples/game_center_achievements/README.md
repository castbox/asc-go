# Game Center Achievements Example

This example demonstrates how to create and configure Game Center achievements using the App Store Connect API.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Verify Credentials](#verify-credentials)
- [Single Achievement Creation](#single-achievement-creation)
- [Batch Achievement Creation](#batch-achievement-creation)
- [Achievement Image Requirements](#achievement-image-requirements)
- [Achievement Ordering with Position Field](#achievement-ordering-with-position-field)
- [Troubleshooting](#troubleshooting)
- [Notes](#notes)

## Prerequisites

1. An App Store Connect API key (Key ID, Issuer ID, and private key file)
2. An app with Game Center enabled or ready to be enabled
3. Achievement images (512x512 PNG format)

## Verify Credentials

Before running the achievement creation scripts, you can verify your API credentials:

```bash
go run verify_credentials.go \
  -kid "YOUR_KEY_ID" \
  -iss "YOUR_ISSUER_ID" \
  -privatekeypath "/path/to/AuthKey_XXXXXX.p8" \
  -bundleid "com.example.yourapp"
```

### Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `-kid` | Yes | App Store Connect API Key ID |
| `-iss` | Yes | App Store Connect Issuer ID |
| `-privatekeypath` | Yes | Path to the private key (.p8 file) |
| `-bundleid` | No | Bundle ID to verify (optional, will list first app if omitted) |

This tool will:
1. Validate your authentication configuration
2. Test the API connection by listing apps
3. Display the first matching app if credentials are valid

## Single Achievement Creation

Create a single achievement with localization and optional image:

```bash
go run main.go \
  -kid "YOUR_KEY_ID" \
  -iss "YOUR_ISSUER_ID" \
  -privatekeypath "/path/to/AuthKey_XXXXXX.p8" \
  -bundleid "com.example.yourapp" \
  -name "First Win" \
  -vendor "com.example.achievement.first_win" \
  -points 10 \
  -showbefore=true \
  -repeatable=false \
  -locale "en-US" \
  -localizedname "First Win" \
  -beforedesc "Win your first game" \
  -afterdesc "You won your first game!" \
  -imagefile "/path/to/achievement_icon.png"
```

### Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `-kid` | Yes | App Store Connect API Key ID |
| `-iss` | Yes | App Store Connect Issuer ID |
| `-privatekeypath` | Yes | Path to the private key (.p8 file) |
| `-bundleid` | Yes | Bundle ID of your app |
| `-name` | Yes | Reference name for the achievement |
| `-vendor` | Yes | Vendor identifier (unique ID for the achievement) |
| `-points` | No | Points for the achievement (1-100, default: 10) |
| `-showbefore` | No | Show achievement before earned (default: true) |
| `-repeatable` | No | Achievement is repeatable (default: false) |
| `-locale` | No | Locale for localization (default: en-US) |
| `-localizedname` | Yes | Localized display name |
| `-beforedesc` | Yes | Description shown before earning |
| `-afterdesc` | Yes | Description shown after earning |
| `-imagefile` | No | Path to achievement image (512x512 PNG) |

## Batch Achievement Creation

Create multiple achievements from a JSON configuration file:

```bash
go run batch_create.go \
  -kid "YOUR_KEY_ID" \
  -iss "YOUR_ISSUER_ID" \
  -privatekeypath "/path/to/AuthKey_XXXXXX.p8" \
  -bundleid "com.example.yourapp" \
  -config "achievements_config.json"
```

### Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `-kid` | Yes | App Store Connect API Key ID |
| `-iss` | Yes | App Store Connect Issuer ID |
| `-privatekeypath` | Yes | Path to the private key (.p8 file) |
| `-bundleid` | Yes | Bundle ID of your app |
| `-config` | Yes | Path to the JSON configuration file |
| `-resume` | No | Resume mode: skip existing achievements and localizations, only upload missing images (default: false) |

### Resume Mode (Incremental Upload)

Use `-resume=true` to enable incremental upload mode:

```bash
go run batch_create.go \
  -kid "YOUR_KEY_ID" \
  -iss "YOUR_ISSUER_ID" \
  -privatekeypath "/path/to/AuthKey_XXXXXX.p8" \
  -bundleid "com.example.yourapp" \
  -config "achievements_config.json" \
  -resume=true
```

In resume mode, the script will:
1. **Skip existing achievements** - If an achievement with the same vendor identifier already exists, it won't be recreated
2. **Skip existing localizations** - If a localization for a specific locale already exists, it won't be recreated
3. **Upload missing images only** - Only upload images for localizations that don't have an image yet

This is useful when:
- A previous batch upload was interrupted
- You want to add new localizations or images to existing achievements
- You want to retry failed image uploads without recreating everything

### Configuration File Format

See `achievements_config.example.json` for a complete example:

```json
{
  "achievements": [
    {
      "referenceName": "First Win",
      "vendorIdentifier": "com.example.achievement.first_win",
      "points": 10,
      "showBeforeEarned": true,
      "repeatable": false,
      "position": 1,
      "localizations": [
        {
          "locale": "en-US",
          "name": "First Win",
          "beforeEarnedDescription": "Win your first game",
          "afterEarnedDescription": "You won your first game!",
          "imageFile": "images/first_win.png"
        },
        {
          "locale": "zh-Hans",
          "name": "首次胜利",
          "beforeEarnedDescription": "赢得你的第一场游戏",
          "afterEarnedDescription": "你赢得了第一场游戏！",
          "imageFile": "images/first_win.png"
        }
      ]
    }
  ]
}
```

## Achievement Image Requirements

- **Size**: 512 x 512 pixels
- **Format**: PNG
- **Color Space**: sRGB or P3
- **Transparency**: Supported

## Notes

1. **Vendor Identifier**: Must be unique across all achievements in your app. Use a reverse-domain style identifier (e.g., `com.yourcompany.game.achievement_name`).

2. **Points**: Must be between 1 and 100. The total points across all achievements should not exceed 1000.

3. **Repeatable Achievements**: Set `repeatable` to `true` for achievements that can be earned multiple times.

4. **Localizations**: You can add multiple localizations for different languages. Each localization can have its own image.

5. **Game Center must be enabled**: The example will automatically enable Game Center for your app if it's not already enabled.

6. **Position (Ordering)**: Each achievement has a `position` field (1-based) that determines its order in the final list.

## Achievement Ordering with Position Field

Each achievement in the config file has a `position` field that controls where it appears in the final order:

| `position` Value | Behavior |
|------------------|----------|
| `1` | Insert at position 1 (first) |
| `2` | Insert at position 2 (second) |
| `N` | Insert at position N |
| `0` or omitted | Append at the end |

### Example: Inserting New Achievements Before Existing Ones

If you already have 3 achievements (A, B, C) and want to add 3 new ones (X, Y, Z) at the beginning:

```json
{
  "achievements": [
    { "referenceName": "X", "position": 1, ... },
    { "referenceName": "Y", "position": 2, ... },
    { "referenceName": "Z", "position": 3, ... }
  ]
}
```

**Result**: X, Y, Z, A, B, C

### Example: Inserting at Specific Positions

If you have achievements (A, B, C, D) and want to insert X at position 2 and Y at position 4:

```json
{
  "achievements": [
    { "referenceName": "X", "position": 2, ... },
    { "referenceName": "Y", "position": 4, ... }
  ]
}
```

**Result**: A, X, B, Y, C, D

### Example: Append at End

Use `position: 0` or omit the field to append at the end:

```json
{
  "achievements": [
    { "referenceName": "New Achievement", "position": 0, ... }
  ]
}
```

**Result**: A, B, C, ..., New Achievement

## Troubleshooting

### Authentication Errors

If you see `401 Unauthorized` or authentication errors:

1. Verify your Key ID and Issuer ID are correct
2. Ensure the private key file path is valid and the file exists
3. Check that the API key has the required permissions in App Store Connect
4. Run `verify_credentials.go` to test your credentials before running other scripts

### Image Upload Errors

- Ensure images are exactly **512x512 pixels**
- Use **PNG format** only
- Check file path is correct (relative to config file directory or absolute path)
- Verify the image file is not corrupted

### Game Center Not Enabled

The script will automatically enable Game Center if needed. If this fails:

1. Manually enable Game Center in App Store Connect
2. Go to your app → App Store → Game Center
3. Enable Game Center and save changes

### API Rate Limiting

If you encounter rate limiting errors:

1. Wait a few minutes before retrying
2. Consider reducing the number of achievements in a single batch
3. The SDK includes automatic retry with exponential backoff
