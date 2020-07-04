# Generated by Django 2.2.9 on 2020-07-04 18:08

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Asset_Element',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=75)),
                ('distance', models.IntegerField()),
                ('satellites', models.ManyToManyField(related_name='element_satellites', to='plan.Asset_Element')),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Asset_Route',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=75)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Asset_Star_Cluster',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=75)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Config_Element',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='-', max_length=25)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Config_Grid',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='-', max_length=25)),
                ('height', models.PositiveSmallIntegerField(blank=True, default=20)),
                ('width', models.PositiveSmallIntegerField(blank=True, default=20)),
                ('connectionRange', models.PositiveSmallIntegerField(blank=True, default=5)),
                ('populationRate', models.FloatField(blank=True, default=0.5)),
                ('connectionRate', models.FloatField(blank=True, default=0.5)),
                ('rangeRateMultiplier', models.FloatField(blank=True, default=0.5)),
                ('smoothingFactor', models.FloatField(blank=True, default=0.5)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Config_Route',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='-', max_length=25)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Config_Star_Cluster',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='-', max_length=25)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Config_System',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='-', max_length=25)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Config_Zone',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='-', max_length=25)),
                ('zone', models.CharField(blank=True, max_length=25, null=True)),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Grid_Sector',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='-', max_length=39)),
            ],
        ),
        migrations.CreateModel(
            name='Inspiration',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=25)),
                ('description', models.CharField(default='', max_length=1000)),
            ],
        ),
        migrations.CreateModel(
            name='Job',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('percent_complete', models.PositiveSmallIntegerField(blank=True, default=0, null=True)),
                ('error', models.TextField(blank=True, default=None, null=True)),
                ('created_at', models.DateTimeField(auto_now_add=True)),
                ('jobType', models.CharField(choices=[('GD', 'Grid'), ('SR', 'Sector')], max_length=2)),
                ('config_id', models.BigIntegerField(null=True)),
                ('asset_id', models.BigIntegerField(blank=True, null=True)),
            ],
        ),
        migrations.CreateModel(
            name='Perterbation',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(default='-', max_length=25)),
                ('element', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, to='plan.Config_Element')),
                ('route', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, to='plan.Config_Route')),
                ('satellite', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, related_name='perterbation_satellites', to='plan.Config_Element')),
                ('star_cluster', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, to='plan.Config_Star_Cluster')),
                ('system', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, to='plan.Config_System')),
            ],
        ),
        migrations.CreateModel(
            name='Roll',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('dice_count', models.PositiveSmallIntegerField(blank=True, default=0)),
                ('dice_size', models.PositiveSmallIntegerField(blank=True, default=6)),
                ('base', models.IntegerField(blank=True, default=0)),
                ('multiplier', models.IntegerField(blank=True, default=1)),
                ('keep_highest', models.IntegerField(blank=True, default=0)),
            ],
        ),
        migrations.CreateModel(
            name='Tag',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=64)),
            ],
        ),
        migrations.CreateModel(
            name='Weighted_Perterbation',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('parent', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='plan.Config_Grid')),
                ('value', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='plan.Perterbation')),
                ('weights', models.ManyToManyField(to='plan.Roll')),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Weighted_Inspiration',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('value', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='plan.Inspiration')),
                ('weights', models.ManyToManyField(to='plan.Roll')),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.AddField(
            model_name='perterbation',
            name='tags',
            field=models.ManyToManyField(to='plan.Tag'),
        ),
        migrations.AddField(
            model_name='perterbation',
            name='zones',
            field=models.ManyToManyField(related_name='zones', to='plan.Config_Zone'),
        ),
        migrations.CreateModel(
            name='Inspiration_Nested',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=25)),
                ('count', models.ManyToManyField(related_name='count', to='plan.Roll')),
                ('weighted_inspirations', models.ManyToManyField(to='plan.Weighted_Inspiration')),
            ],
        ),
        migrations.AddField(
            model_name='inspiration',
            name='nested_inspirations',
            field=models.ManyToManyField(related_name='nested_inspirations', to='plan.Inspiration_Nested'),
        ),
        migrations.AddField(
            model_name='inspiration',
            name='perterbation',
            field=models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.CASCADE, to='plan.Perterbation'),
        ),
        migrations.AddField(
            model_name='inspiration',
            name='roll_groups',
            field=models.ManyToManyField(related_name='roll_groups', to='plan.Inspiration_Nested'),
        ),
        migrations.AddField(
            model_name='inspiration',
            name='tags',
            field=models.ManyToManyField(to='plan.Tag'),
        ),
        migrations.CreateModel(
            name='Grid_System',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('x', models.PositiveSmallIntegerField()),
                ('y', models.PositiveSmallIntegerField()),
                ('region', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='plan.Perterbation')),
                ('sector', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='plan.Grid_Sector')),
            ],
        ),
        migrations.CreateModel(
            name='Grid_Route',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('end', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='end', to='plan.Grid_System')),
                ('start', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='start', to='plan.Grid_System')),
            ],
        ),
        migrations.CreateModel(
            name='Detail',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('rolls', models.CharField(max_length=100)),
                ('inspirations', models.ManyToManyField(related_name='inspirations', to='plan.Inspiration')),
                ('nested_inspirations', models.ManyToManyField(to='plan.Inspiration_Nested')),
                ('parent_detail', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.CASCADE, to='plan.Detail')),
            ],
        ),
        migrations.AddField(
            model_name='config_zone',
            name='distance',
            field=models.ManyToManyField(related_name='distance', to='plan.Roll'),
        ),
        migrations.AddField(
            model_name='config_zone',
            name='element_count',
            field=models.ManyToManyField(related_name='element_count', to='plan.Roll'),
        ),
        migrations.AddField(
            model_name='config_zone',
            name='element_extra',
            field=models.ManyToManyField(related_name='element_extra', to='plan.Weighted_Inspiration'),
        ),
        migrations.AddField(
            model_name='config_zone',
            name='perterbation',
            field=models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.CASCADE, to='plan.Perterbation'),
        ),
        migrations.AddField(
            model_name='config_system',
            name='star_cluster_count',
            field=models.ManyToManyField(related_name='star_cluster_count', to='plan.Roll'),
        ),
        migrations.AddField(
            model_name='config_system',
            name='system_feature_count',
            field=models.ManyToManyField(related_name='system_feature_count', to='plan.Roll'),
        ),
        migrations.AddField(
            model_name='config_system',
            name='system_feature_extra',
            field=models.ManyToManyField(related_name='system_feature_extra', to='plan.Weighted_Inspiration'),
        ),
        migrations.AddField(
            model_name='config_system',
            name='system_feature_inspirations',
            field=models.ManyToManyField(related_name='system_feature_inspirations', to='plan.Weighted_Inspiration'),
        ),
        migrations.AddField(
            model_name='config_star_cluster',
            name='star_count',
            field=models.ManyToManyField(related_name='star_count', to='plan.Roll'),
        ),
        migrations.AddField(
            model_name='config_star_cluster',
            name='star_extra',
            field=models.ManyToManyField(related_name='star_extra', to='plan.Weighted_Inspiration'),
        ),
        migrations.AddField(
            model_name='config_star_cluster',
            name='star_inspirations',
            field=models.ManyToManyField(related_name='star_inspirations', to='plan.Weighted_Inspiration'),
        ),
        migrations.AddField(
            model_name='config_route',
            name='days_inspirations',
            field=models.ManyToManyField(related_name='days_inspirations', to='plan.Weighted_Inspiration'),
        ),
        migrations.AddField(
            model_name='config_route',
            name='stability_inspirations',
            field=models.ManyToManyField(related_name='stability_inspirations', to='plan.Weighted_Inspiration'),
        ),
        migrations.AddField(
            model_name='config_element',
            name='satellite_count',
            field=models.ManyToManyField(related_name='satellite_count', to='plan.Roll'),
        ),
        migrations.AddField(
            model_name='config_element',
            name='satellite_extra',
            field=models.ManyToManyField(related_name='satellite_extra', to='plan.Weighted_Inspiration'),
        ),
        migrations.AddField(
            model_name='config_element',
            name='spacing',
            field=models.ManyToManyField(related_name='spacing', to='plan.Roll'),
        ),
        migrations.AddField(
            model_name='config_element',
            name='type_inspirations',
            field=models.ManyToManyField(related_name='type_inspirations', to='plan.Weighted_Inspiration'),
        ),
        migrations.CreateModel(
            name='Asset_Zone',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=75)),
                ('distance', models.SmallIntegerField()),
                ('elements', models.ManyToManyField(related_name='elements', to='plan.Asset_Element')),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.CreateModel(
            name='Asset_System',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=75)),
                ('details', models.ManyToManyField(related_name='details', to='plan.Detail')),
                ('routes', models.ManyToManyField(related_name='routes', to='plan.Asset_Route')),
                ('star_clusters', models.ManyToManyField(related_name='star_clusters', to='plan.Asset_Star_Cluster')),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.AddField(
            model_name='asset_star_cluster',
            name='stars',
            field=models.ManyToManyField(related_name='stars', to='plan.Detail'),
        ),
        migrations.AddField(
            model_name='asset_star_cluster',
            name='zones',
            field=models.ManyToManyField(related_name='zones', to='plan.Asset_Zone'),
        ),
        migrations.CreateModel(
            name='Asset_Sector',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=75)),
                ('systems', models.ManyToManyField(related_name='systems', to='plan.Asset_System')),
            ],
            options={
                'abstract': False,
            },
        ),
        migrations.AddField(
            model_name='asset_route',
            name='days',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='days', to='plan.Detail'),
        ),
        migrations.AddField(
            model_name='asset_route',
            name='stability',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='stability', to='plan.Detail'),
        ),
        migrations.AddField(
            model_name='asset_route',
            name='target_systems',
            field=models.ManyToManyField(related_name='target_systems', to='plan.Asset_System'),
        ),
        migrations.AddField(
            model_name='asset_element',
            name='type',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='type', to='plan.Detail'),
        ),
    ]
