package com.example.immotep.ui.icons

import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.PathFillType
import androidx.compose.ui.graphics.SolidColor
import androidx.compose.ui.graphics.StrokeCap
import androidx.compose.ui.graphics.StrokeJoin
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.graphics.vector.path
import androidx.compose.ui.unit.dp

public val ReleaseAlert: ImageVector
    get() {
        if (_releaseAlert != null) {
            return _releaseAlert!!
        }
        _releaseAlert =
            ImageVector
                .Builder(
                    name = "Release_alert",
                    defaultWidth = 24.dp,
                    defaultHeight = 24.dp,
                    viewportWidth = 960f,
                    viewportHeight = 960f,
                ).apply {
                    path(
                        fill = SolidColor(Color.White),
                        fillAlpha = 1.0f,
                        stroke = null,
                        strokeAlpha = 1.0f,
                        strokeLineWidth = 1.0f,
                        strokeLineCap = StrokeCap.Butt,
                        strokeLineJoin = StrokeJoin.Miter,
                        strokeLineMiter = 1.0f,
                        pathFillType = PathFillType.NonZero,
                    ) {
                        moveTo(344f, 900f)
                        lineToRelative(-76f, -128f)
                        lineToRelative(-144f, -32f)
                        lineToRelative(14f, -148f)
                        lineToRelative(-98f, -112f)
                        lineToRelative(98f, -112f)
                        lineToRelative(-14f, -148f)
                        lineToRelative(144f, -32f)
                        lineToRelative(76f, -128f)
                        lineToRelative(136f, 58f)
                        lineToRelative(136f, -58f)
                        lineToRelative(76f, 128f)
                        lineToRelative(144f, 32f)
                        lineToRelative(-14f, 148f)
                        lineToRelative(98f, 112f)
                        lineToRelative(-98f, 112f)
                        lineToRelative(14f, 148f)
                        lineToRelative(-144f, 32f)
                        lineToRelative(-76f, 128f)
                        lineToRelative(-136f, -58f)
                        close()
                        moveToRelative(34f, -102f)
                        lineToRelative(102f, -44f)
                        lineToRelative(104f, 44f)
                        lineToRelative(56f, -96f)
                        lineToRelative(110f, -26f)
                        lineToRelative(-10f, -112f)
                        lineToRelative(74f, -84f)
                        lineToRelative(-74f, -86f)
                        lineToRelative(10f, -112f)
                        lineToRelative(-110f, -24f)
                        lineToRelative(-58f, -96f)
                        lineToRelative(-102f, 44f)
                        lineToRelative(-104f, -44f)
                        lineToRelative(-56f, 96f)
                        lineToRelative(-110f, 24f)
                        lineToRelative(10f, 112f)
                        lineToRelative(-74f, 86f)
                        lineToRelative(74f, 84f)
                        lineToRelative(-10f, 114f)
                        lineToRelative(110f, 24f)
                        close()
                        moveToRelative(102f, -118f)
                        quadToRelative(17f, 0f, 28.5f, -11.5f)
                        reflectiveQuadTo(520f, 640f)
                        reflectiveQuadToRelative(-11.5f, -28.5f)
                        reflectiveQuadTo(480f, 600f)
                        reflectiveQuadToRelative(-28.5f, 11.5f)
                        reflectiveQuadTo(440f, 640f)
                        reflectiveQuadToRelative(11.5f, 28.5f)
                        reflectiveQuadTo(480f, 680f)
                        moveToRelative(-40f, -160f)
                        horizontalLineToRelative(80f)
                        verticalLineToRelative(-240f)
                        horizontalLineToRelative(-80f)
                        close()
                    }
                }.build()
        return _releaseAlert!!
    }

private var _releaseAlert: ImageVector? = null

